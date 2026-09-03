function readStoredUser() {
  const rawUser = sessionStorage.getItem('yalix_user') || localStorage.getItem('yalix_user');
  if (!rawUser) return null;
  try { return JSON.parse(rawUser); } catch (e) { return null; }
}

function saveStoredUser(user, remember) {
  const storage = remember ? localStorage : sessionStorage;
  storage.setItem('yalix_user', JSON.stringify(user));
}

function saveToken(token, remember) {
  const storage = remember ? localStorage : sessionStorage;
  storage.setItem('yalix_token', token);
}

const form = document.getElementById('editProfileForm');
const firstNameInput = document.getElementById('firstName');
const lastNameInput = document.getElementById('lastName');
const loginInput = document.getElementById('login');
const emailInput = document.getElementById('email');

const storedUser = readStoredUser();
if (!storedUser) {
  // not logged in, navigate to login
  window.location.replace('login.html');
} else {
  // prefill
  firstNameInput.value = storedUser.first_name || storedUser.firstName || storedUser.name?.split(' ')[0] || '';
  lastNameInput.value = storedUser.last_name || storedUser.lastName || (storedUser.name ? storedUser.name.split(' ').slice(1).join(' ') : '');
  loginInput.value = storedUser.login || storedUser.username || '';
  emailInput.value = storedUser.email || '';
}

form.addEventListener('submit', async (e) => {
  e.preventDefault();
  const updated = {
    ...storedUser,
    first_name: firstNameInput.value.trim(),
    last_name: lastNameInput.value.trim(),
    login: loginInput.value.trim(),
    email: emailInput.value.trim()
  };

  // Save locally first
  // preserve token storage location
  const token = sessionStorage.getItem('yalix_token') || localStorage.getItem('yalix_token');
  const remember = !!localStorage.getItem('yalix_token');
  saveStoredUser(updated, remember);

  // Try to update on server if possible
  try {
    if (typeof apiRequest === 'function') {
      const id = updated.id || updated.user_id || storedUser.id;
      if (id) {
        // Attempt PUT then PATCH as fallback
        try {
          await apiRequest(`/users/${id}`, { method: 'PUT', body: JSON.stringify(updated) });
        } catch (err) {
          // try PATCH
          try {
            await apiRequest(`/users/${id}`, { method: 'PATCH', body: JSON.stringify(updated) });
          } catch (err2) {
            // server may not support update endpoint, ignore
            console.warn('Server update failed', err2);
          }
        }
      } else {
        // No id known, attempt to POST to /users to update? skip.
        console.warn('No user id, skipping server update');
      }
    }

    alert('Данные сохранены.');
    window.location.href = 'personal-cabinet.html';
  } catch (error) {
    console.error('Ошибка при сохранении профиля:', error);
    alert('Локальные данные сохранены, но обновление на сервере не удалось.');
    window.location.href = 'personal-cabinet.html';
  }
});