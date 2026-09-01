const form = document.getElementById('loginForm');
const identifier = document.getElementById('loginIdentifier') || document.getElementById('loginEmail');
const pass = document.getElementById('loginPassword');

if (form && identifier && pass) {
  form.addEventListener('submit', async (e) => {
    e.preventDefault();

    const loginValue = identifier.value.trim();
    const pw = pass.value;

    if (!loginValue || !pw) {
      alert('Введите логин и пароль.');
      return;
    }

    try {
      const payload = {
        login: loginValue,
        username: loginValue,
        email: loginValue,
        password: pw
      };

      const data = await apiRequest('/login', {
        method: 'POST',
        body: JSON.stringify(payload)
      });

      const responseUser = data && typeof data === 'object' ? data.user || data.profile || data : null;
      const responseLogin = responseUser?.login || responseUser?.username || responseUser?.name || loginValue;

      alert(`Вход выполнен успешно для ${responseLogin}`);
      form.reset();
      window.location.href = 'people.html';
    } catch (error) {
      console.error('Login error:', error);
      const message = error?.serverMessage || error?.message || 'неизвестная ошибка';
      alert(`Ошибка входа: ${message}`);
    }
  });
}
