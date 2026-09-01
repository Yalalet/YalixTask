const form = document.getElementById('taskForm');
const formTitle = document.getElementById('formTitle');
const createBtn = document.getElementById('createBtn');
const titleInput = document.getElementById('title');
const descriptionInput = document.getElementById('description');
const dateInput = document.getElementById('dueDate');
const timeInput = document.getElementById('dueTime');
const list = document.getElementById('tasks');
const modal = document.getElementById('taskModal');
const modalTitle = document.getElementById('modalTitle');
const modalMeta = document.getElementById('modalMeta');
const modalDescription = document.getElementById('modalDescription');
const deleteTaskBtn = document.getElementById('deleteTaskBtn');
const closeModalBtns = [document.getElementById('closeModalBtn'), document.getElementById('cancelModalBtn')];
const editTaskBtn = document.getElementById('editTaskBtn');

let activeTaskItem = null;
let editingTaskItem = null;

function resetFormMode() {
  editingTaskItem = null;
  formTitle.textContent = 'Создать задачу';
  createBtn.textContent = 'Создать задачу';
  form.reset();
}

function closeModal() {
  modal.classList.remove('open');
  modal.setAttribute('aria-hidden', 'true');
  activeTaskItem = null;
  document.body.classList.remove('modal-open');
}

function renderTask(taskItem) {
  const title = taskItem.dataset.title || 'Без названия';
  const dueDate = taskItem.dataset.date;
  const dueTime = taskItem.dataset.time;

  const titleEl = taskItem.querySelector('strong');
  const metaEl = taskItem.querySelector('.task-meta');

  titleEl.textContent = title;

  const details = [];
  if (dueDate) details.push(new Date(dueDate + 'T00:00:00').toLocaleDateString('ru-RU'));
  if (dueTime) details.push(dueTime);
  metaEl.textContent = details.join(' • ') || 'Без даты';

  taskItem.dataset.title = title;
  taskItem.dataset.date = dueDate || '';
  taskItem.dataset.time = dueTime || '';
}

function openModal(taskItem) {
  const title = taskItem.dataset.title || taskItem.querySelector('strong')?.textContent || 'Без названия';
  const description = taskItem.dataset.description || 'Описание отсутствует.';
  const dueDate = taskItem.dataset.date;
  const dueTime = taskItem.dataset.time;

  modalTitle.textContent = title;
  const metaParts = [];
  if (dueDate) metaParts.push(new Date(dueDate + 'T00:00:00').toLocaleDateString('ru-RU'));
  if (dueTime) metaParts.push(dueTime);
  modalMeta.textContent = metaParts.length ? metaParts.join(' • ') : 'Дата и время не указаны';
  modalDescription.textContent = description;

  list.querySelectorAll('.task-item').forEach((item) => item.classList.remove('active'));
  taskItem.classList.add('active');
  activeTaskItem = taskItem;

  modal.classList.add('open');
  modal.setAttribute('aria-hidden', 'false');
  document.body.classList.add('modal-open');
}

form.addEventListener('submit', (event) => {
  event.preventDefault();

  const title = titleInput.value.trim();
  const description = descriptionInput.value.trim();
  const dueDate = dateInput.value;
  const dueTime = timeInput.value;

  if (!title) {
    titleInput.classList.add('error');
    titleInput.focus();
    return;
  }

  titleInput.classList.remove('error');

  if (editingTaskItem) {
    editingTaskItem.dataset.title = title;
    editingTaskItem.dataset.description = description || 'Описание отсутствует.';
    editingTaskItem.dataset.date = dueDate || '';
    editingTaskItem.dataset.time = dueTime || '';
    renderTask(editingTaskItem);
    openModal(editingTaskItem);
  } else {
    const li = document.createElement('li');
    li.className = 'task-item';
    li.tabIndex = 0;
    li.setAttribute('role', 'button');
    li.dataset.title = title;
    li.dataset.description = description || 'Описание отсутствует.';
    li.dataset.date = dueDate || '';
    li.dataset.time = dueTime || '';

    const titleEl = document.createElement('strong');
    titleEl.textContent = title;

    const metaEl = document.createElement('div');
    metaEl.className = 'task-meta';

    const details = [];
    if (dueDate) details.push(new Date(dueDate + 'T00:00:00').toLocaleDateString('ru-RU'));
    if (dueTime) details.push(dueTime);
    metaEl.textContent = details.join(' • ') || 'Без даты';

    li.appendChild(titleEl);
    li.appendChild(metaEl);
    list.appendChild(li);

    li.addEventListener('click', () => openModal(li));
    li.addEventListener('keydown', (event) => {
      if (event.key === 'Enter' || event.key === ' ') {
        event.preventDefault();
        openModal(li);
      }
    });
  }

  resetFormMode();
  form.reset();
  titleInput.focus();
});

closeModalBtns.forEach((btn) => btn.addEventListener('click', closeModal));

modal.addEventListener('click', (event) => {
  if (event.target === modal) closeModal();
});

document.addEventListener('keydown', (event) => {
  if (event.key === 'Escape' && modal.classList.contains('open')) {
    closeModal();
  }
});

editTaskBtn.addEventListener('click', () => {
  if (!activeTaskItem) return;

  editingTaskItem = activeTaskItem;
  formTitle.textContent = 'Редактировать задачу';
  createBtn.textContent = 'Сохранить изменения';

  const title = activeTaskItem.dataset.title || '';
  const description = activeTaskItem.dataset.description || '';
  const dueDate = activeTaskItem.dataset.date || '';
  const dueTime = activeTaskItem.dataset.time || '';

  titleInput.value = title;
  descriptionInput.value = description === 'Описание отсутствует.' ? '' : description;
  dateInput.value = dueDate;
  timeInput.value = dueTime;

  closeModal();
  titleInput.focus();
  window.scrollTo({ top: 0, behavior: 'smooth' });
});

deleteTaskBtn.addEventListener('click', () => {
  if (!activeTaskItem) return;

  const confirmed = window.confirm('Удалить эту задачу?');
  if (!confirmed) return;

  activeTaskItem.remove();
  closeModal();
});
