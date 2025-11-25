import React, { useState, useEffect } from 'react';
import AddCourseForm from './components/AddCourseForm';
import CourseCard from './components/coursecard'; // ← Убедитесь, что имя файла совпадает!

const API_URL = '/api/courses';

function App() {
  const [courses, setCourses] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');
  const [editingCourse, setEditingCourse] = useState(null);


  const openEditModal = (course) => {
    setEditingCourse({ ...course }); // копируем, чтобы не мутировать оригинал
  };

  const closeEditModal = () => setEditingCourse(null);



  // Загрузка курсов
  const loadCourses = async () => {
    try {
      const res = await fetch(API_URL);
      if (!res.ok) throw new Error('Failed to load courses');
      const data = await res.json();
      setCourses(data);
    } catch (err) {
      setError(err.message);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => { loadCourses(); }, []);

  // Добавление курса — передаём данные в форму через callback
  const handleAddCourse = async (formData) => {
    try {
      const res = await fetch(API_URL, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(formData),
      });

      if (!res.ok) {
        const err = await res.json();
        throw new Error(err.error || 'Ошибка создания курса');
      }

      const newCourse = await res.json();
      setCourses([newCourse, ...courses]);
      return true; // Успешно

    } catch (err) {
      console.error('Ошибка при добавлении курса:', err);
      return false;
    }
  };
  // Обновление курса
const handleUpdateCourse = async (id, updatedData) => {
  try {
    const res = await fetch(`${API_URL}/${id}`, {
      method: 'PUT', // или 'PATCH', как у вас в Go
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(updatedData),
    });

    if (!res.ok) {
      const err = await res.json();
      throw new Error(err.error || 'Не удалось обновить курс');
    }

    const updatedCourse = await res.json();
    setCourses(prev =>
      prev.map(course => course.id === id ? updatedCourse : course)
    );
    return true;
  } catch (err) {
    console.error('Ошибка обновления:', err);
    alert('Ошибка: ' + err.message);
    return false;
  }
  };
  // Удаление курса
  const handleDeleteCourse = async (id) => {
    if (!window.confirm('Вы уверены, что хотите удалить курс?')) return;

    try {
      const res = await fetch(`${API_URL}/${id}`, { method: 'DELETE' });
      if (!res.ok) {
        const err = await res.json();
        throw new Error(err.error || 'Ошибка удаления');
      }

      setCourses(courses.filter(c => c.id !== id));
    } catch (err) {
      alert('Ошибка: ' + err.message);
    }
  };

  if (loading) return (
    <div className="min-h-screen flex items-center justify-center bg-gradient-to-br from-indigo-50 to-blue-100">
      <div className="text-center">
        <div className="inline-block animate-spin rounded-full h-12 w-12 border-t-2 border-b-2 border-indigo-600 mb-4"></div>
        <p className="text-gray-700 font-medium">Загрузка курсов...</p>
      </div>
    </div>
  );

  if (error) return (
    <div className="min-h-screen flex items-center justify-center bg-gradient-to-br from-red-50 to-orange-50 p-4">
      <div className="bg-white rounded-xl shadow-lg p-6 max-w-md text-center">
        <div className="text-red-500 text-2xl mb-2">⚠️</div>
        <h2 className="text-xl font-bold text-gray-800 mb-2">Ошибка загрузки</h2>
        <p className="text-gray-600 mb-4">{error}</p>
        <button
          onClick={loadCourses}
          className="px-4 py-2 bg-indigo-600 text-white rounded-md hover:bg-indigo-700 transition-colors"
        >
          Повторить
        </button>
      </div>
    </div>
  );

  return (
    <div className="min-h-screen bg-gradient-to-br from-indigo-50 to-blue-100 p-4">
      <div className="max-w-4xl mx-auto">

        {/* Заголовок */}
        <header className="text-center mb-10 pt-6">
          <h1 className="text-4xl font-extrabold text-transparent bg-clip-text bg-gradient-to-r from-indigo-600 to-purple-600">
            🎓 Менеджер курсов
          </h1>
          <p className="text-gray-600 mt-2">Создавайте, просматривайте и управляйте своими курсами</p>
        </header>

        {/* Форма добавления — используем компонент */}
        <section className="bg-white rounded-2xl shadow-xl p-6 mb-10 border border-gray-200">
          <h2 className="text-2xl font-bold text-gray-800 mb-6 pb-2 border-b border-gray-100">
            ➕ Добавить новый курс
          </h2>
          <AddCourseForm onAdd={handleAddCourse} />
        </section>

        {/* Список курсов — используем компонент CourseCard */}
        <section>
          <div className="flex justify-between items-center mb-6">
            <h2 className="text-2xl font-bold text-gray-800">
              📚 Ваши курсы ({courses.length})
            </h2>
          </div>

          {courses.length === 0 ? (
            <div className="text-center py-12 bg-white rounded-2xl shadow-lg border border-dashed border-gray-300">
              <div className="text-5xl mb-4">📭</div>
              <h3 className="text-xl font-semibold text-gray-700 mb-2">Пока нет курсов</h3>
              <p className="text-gray-500 max-w-md mx-auto">
                Нажмите кнопку выше, чтобы создать первый курс — и начните обучать мир!
              </p>
            </div>
          ) : (
            <div className="grid gap-6 md:grid-cols-2 lg:grid-cols-3">
              {courses.map(course => (
                <CourseCard
                  key={course.id}
                  course={course}
                  onDelete={handleDeleteCourse}
                  onEdit={openEditModal} 
                />
              ))}
            </div>
          )}
        </section>
      </div>
      {/* --- МОДАЛЬНОЕ ОКНО РЕДАКТИРОВАНИЯ --- */}
    {editingCourse && (
      <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 p-4">
        <div className="bg-white rounded-2xl shadow-2xl w-full max-w-md">
          <div className="p-6">
            <div className="flex justify-between items-center mb-4">
              <h3 className="text-xl font-bold text-gray-800">Редактировать курс</h3>
              <button
                onClick={closeEditModal}
                className="text-gray-500 hover:text-gray-700 text-2xl"
              >
                &times;
              </button>
            </div>

            <form
              onSubmit={async (e) => {
                e.preventDefault();
                const updated = {
                  title: editingCourse.title,
                  description: editingCourse.description,
                  author: editingCourse.author,
                };
                const success = await handleUpdateCourse(editingCourse.id, updated);
                if (success) closeEditModal();
              }}
              className="space-y-4"
            >
              <div>
                <label className="block text-sm font-medium text-gray-700 mb-1">
                  Название *
                </label>
                <input
                  value={editingCourse.title}
                  onChange={(e) => setEditingCourse({ ...editingCourse, title: e.target.value })}
                  className="w-full px-3 py-2 border border-gray-300 rounded-md focus:ring-2 focus:ring-indigo-500"
                  required
                />
              </div>

              <div>
                <label className="block text-sm font-medium text-gray-700 mb-1">
                  Описание
                </label>
                <textarea
                  value={editingCourse.description}
                  onChange={(e) => setEditingCourse({ ...editingCourse, description: e.target.value })}
                  rows="3"
                  className="w-full px-3 py-2 border border-gray-300 rounded-md focus:ring-2 focus:ring-indigo-500"
                />
              </div>

              <div>
                <label className="block text-sm font-medium text-gray-700 mb-1">
                  Автор *
                </label>
                <input
                  value={editingCourse.author}
                  onChange={(e) => setEditingCourse({ ...editingCourse, author: e.target.value })}
                  className="w-full px-3 py-2 border border-gray-300 rounded-md focus:ring-2 focus:ring-indigo-500"
                  required
                />
              </div>

              <div className="flex space-x-3 pt-2">
                <button
                  type="submit"
                  className="flex-1 py-2 px-4 bg-indigo-600 text-white rounded-lg hover:bg-indigo-700"
                >
                  Сохранить
                </button>
                <button
                  type="button"
                  onClick={closeEditModal}
                  className="flex-1 py-2 px-4 bg-gray-200 text-gray-800 rounded-lg hover:bg-gray-300"
                >
                  Отмена
                </button>
              </div>
            </form>
          </div>
        </div>
      </div>
      )}
{/* --- КОНЕЦ МОДАЛЬНОГО ОКНА --- */}
    </div>
  );
}

export default App;


// import React, { useState, useEffect } from 'react';
// import AddCourseForm from './components/AddCourseForm';
// import CourseCard from './components/coursecard';

// const API_URL = '/api/courses';

// function App() {
//   const [courses, setCourses] = useState([]);
//   const [loading, setLoading] = useState(true);
//   const [error, setError] = useState('');

//   // Форма добавления
//   const [title, setTitle] = useState('');
//   const [description, setDescription] = useState('');
//   const [author, setAuthor] = useState('');
//   const [formError, setFormError] = useState('');

//   // Загрузка курсов
//   const loadCourses = async () => {
//     try {
//       const res = await fetch(API_URL);
//       if (!res.ok) throw new Error('Failed to load courses');
//       const data = await res.json();
//       setCourses(data);
//     } catch (err) {
//       setError(err.message);
//     } finally {
//       setLoading(false);
//     }
//   };

//   useEffect(() => { loadCourses(); }, []);

//   // Добавление курса
//   const handleCreate = async (e) => {
//     e.preventDefault();
//     setFormError('');

//     // Валидация фронтенда
//     if (!title.trim()) {
//       setFormError('Название обязательно');
//       return;
//     }
//     if (title.trim().length < 3) {
//       setFormError('Название должно быть не короче 3 символов');
//       return;
//     }
//     if (!author.trim()) {
//       setFormError('Автор обязателен');
//       return;
//     }

//     try {
//       const res = await fetch(API_URL, {
//         method: 'POST',
//         headers: { 'Content-Type': 'application/json' },
//         body: JSON.stringify({ title, description, author }),
//       });

//       if (!res.ok) {
//         const err = await res.json();
//         throw new Error(err.error || 'Ошибка создания курса');
//       }

//       const newCourse = await res.json();
//       setCourses([newCourse, ...courses]);

//       // Очистка формы
//       setTitle('');
//       setDescription('');
//       setAuthor('');

//     } catch (err) {
//       setFormError(err.message);
//     }
//   };

//   // Удаление курса
//   const handleDelete = async (id) => {
//     if (!window.confirm('Вы уверены, что хотите удалить курс?')) return;

//     try {
//       const res = await fetch(`${API_URL}/${id}`, { method: 'DELETE' });
//       if (!res.ok) {
//         const err = await res.json();
//         throw new Error(err.error || 'Ошибка удаления');
//       }

//       setCourses(courses.filter(c => c.id !== id));
//     } catch (err) {
//       alert('Ошибка: ' + err.message);
//     }
//   };

//   if (loading) return <div className="container">Загрузка...</div>;
//   if (error) return <div className="container error">Ошибка: {error}</div>;

//   return (
//     <div className="container">
//       <h1>Менеджер курсов</h1>

//       {/* Форма добавления */}
//       <form onSubmit={handleCreate} className="form">
//         <h2>Добавить курс</h2>
//         {formError && <div className="error">{formError}</div>}

//         <div className="form-group">
//           <label>Название *</label>
//           <input
//             value={title}
//             onChange={(e) => setTitle(e.target.value)}
//             placeholder="Например, 'Основы Go'"
//             required
//           />
//         </div>

//         <div className="form-group">
//           <label>Описание</label>
//           <textarea
//             value={description}
//             onChange={(e) => setDescription(e.target.value)}
//             placeholder="Краткое описание курса"
//           />
//         </div>

//         <div className="form-group">
//           <label>Автор *</label>
//           <input
//             value={author}
//             onChange={(e) => setAuthor(e.target.value)}
//             placeholder="Ваше имя"
//             required
//           />
//         </div>

//         <button type="submit" className="btn btn-primary">
//           Добавить курс
//         </button>
//       </form>
      
//       {/* Список курсов */}
//       <div className="courses">
//         <h2>Список курсов ({courses.length})</h2>
//         {courses.length === 0 ? (
//           <p>Нет курсов. Добавьте первый!</p>
//         ) : (
//           courses.map(course => (
//             <div key={course.id} className="course-card">
//               <h3>{course.title}</h3>
//               <p>{course.description || <em>Без описания</em>}</p>
//               <p className="author">Автор: {course.author}</p>
//               <p className="date">
//                 Создан: {new Date(course.created_at).toLocaleString('ru-RU')}
//               </p>
//               <button
//                 onClick={() => handleDelete(course.id)}
//                 className="btn btn-danger"
//               >
//                 Удалить
//               </button>
//             </div>
//           ))
//         )}
//       </div>
//     </div>
//   );
// }

// export default App;
