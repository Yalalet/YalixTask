const form = document.getElementById('registerForm');
const firstNameInput = document.getElementById('regFirstName');
const lastNameInput = document.getElementById('regLastName');
const loginInput = document.getElementById('regLogin');
const emailInput = document.getElementById('regEmail');
const passInput = document.getElementById('regPassword');
const pass2Input = document.getElementById('regPassword2');

if (form && firstNameInput && lastNameInput && loginInput && emailInput && passInput && pass2Input) {
  form.addEventListener('submit', async (e) => {
    e.preventDefault();

    const firstName = firstNameInput.value.trim();
    const lastName = lastNameInput.value.trim();
    const login = loginInput.value.trim();
    const email = emailInput.value.trim();
    const p1 = passInput.value;
    const p2 = pass2Input.value;

    if (!firstName || !lastName || !login || !email) {
      alert('Заполните имя, фамилию, логин и почту.');
      return;
    }
    if (p1.length < 6) {
      alert('Пароль должен быть минимум 6 символов.');
      return;
    }
    if (p1 !== p2) {
      alert('Пароли не совпадают.');
      return;
    }

    const fullName = `${firstName} ${lastName}`.trim();

    try {
      const payload = {
        first_name: firstName,
        last_name: lastName,
        login,
        name: fullName,
        email,
        password: p1
      };

      const data = await apiRequest('/users', {
        method: 'POST',
        body: JSON.stringify(payload)
      });

      const responseUser = data && typeof data === 'object' ? data : null;
      const responseEmail = responseUser?.email || responseUser?.user_email || email;

      alert(`Аккаунт успешно создан для ${responseEmail}.`);
      form.reset();
      window.location.href = 'login.html';
    } catch (error) {
      console.error('Registration error:', error);
      const message = error?.serverMessage || error?.message || 'неизвестная ошибка';
      alert(`Ошибка регистрации: ${message}`);
    }
  });
}
