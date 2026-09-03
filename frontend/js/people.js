const peopleList = document.getElementById('peopleList');
const peopleLoading = document.getElementById('peopleLoading');
const peopleError = document.getElementById('peopleError');
const refreshPeopleBtn = document.getElementById('refreshPeopleBtn');

function showError(message) {
  if (!peopleLoading || !peopleError) return;

  peopleLoading.classList.add('hidden');
  peopleError.textContent = message;
  peopleError.classList.remove('hidden');
}

function renderPeople(users) {
  if (!peopleList || !peopleLoading || !peopleError) return;

  peopleList.innerHTML = '';

  if (!Array.isArray(users) || users.length === 0) {
    const emptyState = document.createElement('li');
    emptyState.className = 'people-empty';
    emptyState.textContent = 'Пользователи не найдены.';
    peopleList.appendChild(emptyState);
    peopleLoading.classList.add('hidden');
    peopleError.classList.add('hidden');
    return;
  }

  users.forEach((user) => {
    const safeUser = normalizeUser(user);
    const li = document.createElement('li');
    li.className = 'person-card';

    const displayName = typeof safeUser.name === 'string' ? safeUser.name : 'Пользователь';
    const initials = displayName
      .split(' ')
      .map((part) => part[0])
      .join('')
      .slice(0, 2)
      .toUpperCase();

    const avatar = document.createElement('div');
    avatar.className = 'person-avatar';
    avatar.textContent = initials || 'U';

    const info = document.createElement('div');
    info.className = 'person-info';

    const name = document.createElement('div');
    name.className = 'person-name';
    name.textContent = displayName || 'Без имени';

    const login = document.createElement('div');
    login.className = 'person-login';
    login.textContent = `@${safeUser.login || safeUser.username || '—'}`;

    const role = document.createElement('div');
    role.className = 'person-role';
    role.textContent = `Роль: ${safeUser.role_name || 'Роль не указана'}`;

    info.append(name, login, role);

    const id = document.createElement('span');
    id.className = 'person-id';
    id.textContent = `#${safeUser.id ?? 'N/A'}`;

    li.append(avatar, info, id);

    peopleList.appendChild(li);
  });

  peopleLoading.classList.add('hidden');
  peopleError.classList.add('hidden');
}

async function loadPeople() {
  if (!peopleList || !peopleLoading || !peopleError) return;

  peopleLoading.classList.remove('hidden');
  peopleError.classList.add('hidden');
  peopleList.innerHTML = '';

  try {
    const data = await fetchUsers();
    renderPeople(data);
  } catch (error) {
    peopleList.innerHTML = '';
    peopleLoading.classList.add('hidden');
    peopleError.textContent = `Не удалось получить пользователей с сервера. ${error.message || 'Проверьте, что API запущен.'}`;
    peopleError.classList.remove('hidden');
  }
}

if (refreshPeopleBtn) {
  refreshPeopleBtn.addEventListener('click', loadPeople);
}

loadPeople();
