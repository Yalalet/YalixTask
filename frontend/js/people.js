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
    peopleList.innerHTML = '<li class="people-empty">Пользователи не найдены.</li>';
    peopleLoading.classList.add('hidden');
    peopleError.classList.add('hidden');
    return;
  }

  users.forEach((user) => {
    const safeUser = normalizeUser(user);
    const li = document.createElement('li');
    li.className = 'person-card';

    const initials = (safeUser.name || 'U')
      .split(' ')
      .map((part) => part[0])
      .join('')
      .slice(0, 2)
      .toUpperCase();

    li.innerHTML = `
      <div class="person-avatar">${initials}</div>
      <div class="person-info">
        <div class="person-name">${safeUser.name || 'Без имени'}</div>
          <div class="person-login">@${safeUser.login || safeUser.username || '—'}</div>
          <div class="person-role">Роль: ${safeUser.role_name || 'Роль не указана'}</div>
      </div>
      <span class="person-id">#${safeUser.id ?? 'N/A'}</span>
    `;

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
