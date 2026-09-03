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
    const emptyState = document.createElement('li');
    emptyState.className = 'people-empty';
    emptyState.textContent = 'Команды не найдены.';
    teamsList.appendChild(emptyState);
    teamsLoading.classList.add('hidden');
    teamsError.classList.add('hidden');
    return;
  }

  teams.forEach((team) => {
    const li = document.createElement('li');
    li.className = 'person-card';

    const teamName = typeof team.name === 'string'
      ? team.name
      : (typeof team.first_name === 'string' ? team.first_name : '');
    const initials = (teamName || 'T')
      .split(' ')
      .map((part) => part[0])
      .join('')
      .slice(0, 2)
      .toUpperCase();

    const membersCount = Array.isArray(team.members) ? team.members.length : (team.members_count ?? (team.size ?? '—'));

    const avatar = document.createElement('div');
    avatar.className = 'person-avatar';
    avatar.textContent = initials || 'T';

    const info = document.createElement('div');
    info.className = 'person-info';

    const name = document.createElement('div');
    name.className = 'person-name';
    name.textContent = teamName || 'Без названия';

    const members = document.createElement('div');
    members.className = 'person-login';
    members.textContent = `Участников: ${membersCount}`;

    info.append(name, members);

    const id = document.createElement('span');
    id.className = 'person-id';
    id.textContent = `#${team.id ?? 'N/A'}`;

    li.append(avatar, info, id);

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
