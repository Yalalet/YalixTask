const teamsList = document.getElementById('teamsList');
const teamsLoading = document.getElementById('teamsLoading');
const teamsError = document.getElementById('teamsError');
const refreshTeamsBtn = document.getElementById('refreshTeamsBtn');

function showTeamsError(message) {
  if (!teamsLoading || !teamsError) return;
  teamsLoading.classList.add('hidden');
  teamsError.textContent = message;
  teamsError.classList.remove('hidden');
}

function renderTeams(teams) {
  if (!teamsList || !teamsLoading || !teamsError) return;

  teamsList.innerHTML = '';

  if (!Array.isArray(teams) || teams.length === 0) {
    teamsList.innerHTML = '<li class="people-empty">Команды не найдены.</li>';
    teamsLoading.classList.add('hidden');
    teamsError.classList.add('hidden');
    return;
  }

  teams.forEach((team) => {
    const li = document.createElement('li');
    li.className = 'person-card';

    const initials = (team.name || 'T')
      .split(' ')
      .map((part) => part[0])
      .join('')
      .slice(0, 2)
      .toUpperCase();

    const membersCount = Array.isArray(team.members) ? team.members.length : (team.members_count ?? (team.size ?? '—'));

    li.innerHTML = `
      <div class="person-avatar">${initials}</div>
      <div class="person-info">
        <div class="person-name">${team.name || 'Без названия'}</div>
        <div class="person-login">Участников: ${membersCount}</div>
      </div>
      <span class="person-id">#${team.id ?? 'N/A'}</span>
    `;

    teamsList.appendChild(li);
  });

  teamsLoading.classList.add('hidden');
  teamsError.classList.add('hidden');
}

async function fetchTeams() {
  const data = await apiRequest('/teams', { method: 'GET' });

  // Normalize response shapes
  if (Array.isArray(data)) return data;
  if (data && Array.isArray(data.teams)) return data.teams;
  if (data && Array.isArray(data.data)) return data.data;
  return [];
}

async function loadTeams() {
  if (!teamsList || !teamsLoading || !teamsError) return;

  teamsLoading.classList.remove('hidden');
  teamsError.classList.add('hidden');
  teamsList.innerHTML = '';

  try {
    const data = await fetchTeams();
    renderTeams(data);
  } catch (error) {
    teamsList.innerHTML = '';
    teamsLoading.classList.add('hidden');
    const message = error?.serverMessage || error?.message || 'Неизвестная ошибка';
    teamsError.textContent = `Не удалось получить команды. ${message}`;
    teamsError.classList.remove('hidden');
  }
}

if (refreshTeamsBtn) {
  refreshTeamsBtn.addEventListener('click', loadTeams);
}

loadTeams();
