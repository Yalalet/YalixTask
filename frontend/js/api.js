const API_BASE_URL_CANDIDATES = (() => {
  const candidates = [];

  try {
    const saved = localStorage.getItem('yalix_api_base');
    if (saved && /^https?:\/\//i.test(saved)) {
      try {
        const u = new URL(saved);
        const savedHost = u.hostname;
        const currentHost = (typeof window !== 'undefined' && window.location) ? window.location.hostname : null;
        const savedIsLoopback = savedHost === '127.0.0.1' || savedHost === 'localhost';
        const currentIsLoopback = currentHost === '127.0.0.1' || currentHost === 'localhost';
        // If saved points to loopback but current access is via LAN IP, prefer dynamic host later — skip saved here.
        if (!savedIsLoopback || currentIsLoopback || !currentHost) {
          candidates.push(saved.replace(/\/$/, ''));
        }
      } catch (e) {
        candidates.push(saved.replace(/\/$/, ''));
      }
    }
  } catch (error) {
    // Ignore storage issues.
  }

  const hostCandidates = [];
  try {
    if (typeof window !== 'undefined' && window.location && window.location.hostname) {
      const host = window.location.hostname;
      hostCandidates.push(`http://${host}:8080`);
      hostCandidates.push(`http://${host}:8000`);
    }
  } catch (e) {
    // ignore
  }

  // Fallback to local loopback hosts for desktop scenarios
  hostCandidates.push('http://127.0.0.1:8080');
  hostCandidates.push('http://localhost:8080');
  hostCandidates.push('http://127.0.0.1:8000');
  hostCandidates.push('http://localhost:8000');

  for (const host of hostCandidates) {
    if (!candidates.includes(host)) {
      candidates.push(host);
    }
  }

  return candidates;
})();

const API_BASE_URL = API_BASE_URL_CANDIDATES[0];

async function parseErrorMessage(response, fallbackMessage) {
  const text = await response.text().catch(() => '');

  if (!text) {
    return fallbackMessage;
  }

  try {
    const parsed = JSON.parse(text);
    return parsed.message || parsed.error || parsed.details || fallbackMessage;
  } catch (error) {
    return text || fallbackMessage;
  }
}

async function apiRequest(path, options = {}) {
  let lastError = null;

  for (const base of API_BASE_URL_CANDIDATES) {
    try {
      const response = await fetch(`${base}${path}`, {
        headers: {
          'Content-Type': 'application/json',
          ...(options.headers || {})
        },
        ...options
      });

      const rawText = await response.text();
      let payload = null;

      if (rawText && rawText.trim()) {
        try {
          payload = JSON.parse(rawText);
        } catch (error) {
          payload = rawText;
        }
      }

      if (!response.ok) {
        const message = payload && (payload.message || payload.error || payload.details)
          ? payload.message || payload.error || payload.details
          : rawText || 'Ошибка сервера';

        const err = new Error(message);
        err.status = response.status;
        err.serverMessage = message;
        // Prefer keeping a previously recorded serverMessage (don't overwrite with later network errors)
        if (!lastError || !lastError.serverMessage) {
          lastError = err;
        }
        continue;
      }

      return payload;
    } catch (error) {
      // Only record network/fetch errors if we don't already have a server-side message
      if (!lastError || !lastError.serverMessage) {
        lastError = error;
      }
    }
  }

  if (lastError) {
    const message = lastError && lastError.serverMessage
      ? lastError.serverMessage
      : lastError && lastError.message
        ? lastError.message
        : 'Не удалось выполнить запрос к API.';

    const isNetworkError = message.includes('Failed to fetch') || message.includes('ERR_CONNECTION_REFUSED') || message.includes('NetworkError');

    const finalMessage = isNetworkError
      ? 'Сервер временно недоступен. Проверьте подключение, запущен ли бэкенд и правильность адреса API. Если сервер работает, проверьте настройки CORS.'
      : message;

    const error = new Error(finalMessage);
    error.serverMessage = lastError && lastError.serverMessage ? lastError.serverMessage : undefined;
    error.status = lastError && lastError.status ? lastError.status : undefined;
    throw error;
  }

  throw new Error('Не удалось выполнить запрос к API.');
}

function normalizeUser(user = {}) {
  const firstName = typeof user.first_name === 'string' ? user.first_name.trim() : '';
  const lastName = typeof user.last_name === 'string' ? user.last_name.trim() : '';
  const rawName = typeof user.name === 'string' ? user.name.trim() : '';
  const resolvedName = rawName || [firstName, lastName].filter(Boolean).join(' ') || user.username || 'Пользователь';

  const roleId = user.role_id ?? user.roleId ?? (user.role && typeof user.role === 'object' ? user.role.id : null);
  const roleName = user.role_name || user.roleName || (user.role && typeof user.role === 'object' ? user.role.name : '') ||
    (roleId === 1 ? 'Администратор' : roleId === 2 ? 'Пользователь' : '');

  return {
    ...user,
    id: user.id ?? user.user_id ?? null,
    login: user.login || user.username || '',
    name: resolvedName,
    first_name: firstName,
    last_name: lastName,
    email: user.email || user.user_email || '',
    role_id: roleId ?? null,
    role_name: roleName || 'Роль не указана'
  };
}

function normalizeUsersResponse(data) {
  if (Array.isArray(data)) {
    return data.map(normalizeUser);
  }

  if (data && Array.isArray(data.users)) {
    return data.users.map(normalizeUser);
  }

  if (data && Array.isArray(data.data)) {
    return data.data.map(normalizeUser);
  }

  return [];
}

async function fetchUsers() {
  const data = await apiRequest('/users', { method: 'GET' });
  return normalizeUsersResponse(data);
}
