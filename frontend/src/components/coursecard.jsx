import React, { useState } from 'react';

const CourseCard = ({ course, onDelete, onEdit }) => {
  return (
    <div className="bg-white rounded-2xl shadow-lg overflow-hidden border border-gray-200 hover:shadow-xl transition-shadow duration-300">
      <div className="p-5">
        <h3 className="text-xl font-bold text-gray-800 mb-2 line-clamp-1">
          {course.title}
        </h3>
        <p className="text-gray-600 mb-4 line-clamp-3 min-h-[60px]">
          {course.description || <em className="text-gray-400">Без описания</em>}
        </p>

        <div className="flex justify-between items-center text-sm text-gray-500 mb-4">
          <span className="bg-indigo-100 text-indigo-800 px-2 py-1 rounded-full">
            👤 {course.author}
          </span>
          <span title={new Date(course.created_at).toLocaleString('ru-RU')}>
            📅 {new Date(course.created_at).toLocaleDateString('ru-RU')}
          </span>
        </div>

        <div className="flex space-x-2">
          <button
            onClick={() => onEdit(course)}
            className="flex-1 py-2 px-3 bg-amber-500 text-white font-medium rounded-lg hover:bg-amber-600 focus:outline-none focus:ring-2 focus:ring-amber-400 transition-colors text-sm"
          >
            ✏️ Редактировать
          </button>
          <button
            onClick={() => onDelete(course.id)}
            className="flex-1 py-2 px-3 bg-red-500 text-white font-medium rounded-lg hover:bg-red-600 focus:outline-none focus:ring-2 focus:ring-red-400 transition-colors text-sm"
          >
            🗑️ Удалить
          </button>
        </div>
      </div>
    </div>
  );
};

export default CourseCard;