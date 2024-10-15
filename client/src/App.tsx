import React, { useCallback, useEffect, useState } from 'react';
import { useForm, SubmitHandler } from 'react-hook-form';

// Define the Todo type
type Task = {
  id?: string;
  title: string;
  content?: string;
  done?: boolean;
};

// Form input type (for creating/updating a todo)
type TaskFormInput = {
  title: string;
  content?: string;
};

// Todo App Component
const TodoApp: React.FC = () => {
  const { register, handleSubmit, reset, setValue } = useForm<TaskFormInput>();
  const [tasks, setTasks] = useState<Task[]>([]);
  const [isEditing, setIsEditing] = useState<boolean>(false);
  const [editIndex, setEditIndex] = useState<number | null>(null);

  const API_BASE_URL = 'http://localhost:8081/api/v1';

  const getListTasks = useCallback(async () => {
    return fetch(`${API_BASE_URL}/tasks`, { method: 'GET' })
      .then((res) => res.json())
      .then((data) => {
        console.log('data', data);

        return data;
      })
      .catch((err) => {
        console.log('err', err);

        return [];
      });
  }, []);

  const createTask = useCallback(async (task: Task) => {
    await fetch(`${API_BASE_URL}/tasks`, {
      method: 'POST',
      body: JSON.stringify(task),
      headers: {
        'Content-Type': 'application/json',
      },
    })
      .then((res) => res.json())
      .then((data) => {
        console.log('data', data);
      })
      .catch((err) => {
        console.log('err', err);

        throw new Error(err);
      });
  }, []);

  const updateTask = useCallback(async (id: string, task: Task) => {
    await fetch(`${API_BASE_URL}/tasks/${id}`, {
      method: 'PUT',
      body: JSON.stringify(task),
      headers: {
        'Content-Type': 'application/json',
      },
    })
      .then((res) => res.json())
      .then((data) => {
        console.log('data', data);
      })
      .catch((err) => {
        console.log('err', err);

        throw new Error(err);
      });
  }, []);

  const handleTaskAsDone = useCallback(async (id: string) => {
    await fetch(`${API_BASE_URL}/tasks/${id}/done`, {
      method: 'PATCH',
      headers: {
        'Content-Type': 'application/json',
      },
    })
      .then((res) => res.json())
      .then((data) => {
        console.log('data', data);
      })
      .catch((err) => {
        console.log('err', err);

        throw new Error(err);
      });
  }, []);

  const handleArchiveTask = useCallback(async (id: string) => {
    await fetch(`${API_BASE_URL}/tasks/${id}`, {
      method: 'DELETE',
      headers: {
        'Content-Type': 'application/json',
      },
    })
      .then((res) => res.json())
      .then((data) => {
        console.log('data', data);
      })
      .catch((err) => {
        console.log('err', err);

        throw new Error(err);
      });
  }, []);

  const loadTasks = useCallback(() => {
    getListTasks().then((data) => setTasks(data));
  }, [getListTasks]);

  useEffect(() => {
    loadTasks();
  }, [loadTasks]);

  // Handle form submission
  const onSubmit: SubmitHandler<TaskFormInput> = (data) => {
    if (isEditing && editIndex !== null) {
      const todo = tasks[editIndex];
      console.log('data', data);
      console.log('todo', todo);
      if (todo) {
        updateTask(todo.id as string, {
          title: data.title,
          content: data.content,
        }).then(() => {
          setIsEditing(false);
          setEditIndex(null);
          loadTasks();
        });
      }
    } else {
      createTask({ title: data.title, content: data.content || '' }).then(
        () => {
          loadTasks();
        }
      );
    }
    reset(); // Reset form after submission
  };

  const handleEditTask = (index: number) => {
    const task = tasks[index];
    setValue('title', task.title);
    setValue('content', task.content || '');
    setIsEditing(true);
    setEditIndex(index);
  };

  const handleMoveToDone = (index: number) => {
    const todo = tasks[index];
    if (todo) {
      handleTaskAsDone(todo.id as string).then(() => {
        loadTasks();
      });
    }
  };

  const handleMoveToArchive = (index: number) => {
    const todo = tasks[index];
    if (todo) {
      handleArchiveTask(todo.id as string).then(() => {
        loadTasks();
      });
    }
  };

  return (
    <div className='min-h-screen bg-gray-100 flex items-center justify-center p-5'>
      <div className='bg-white p-6 rounded-lg shadow-lg w-full max-w-lg'>
        <h1 className='text-2xl font-bold text-gray-800 mb-4'>Todo List</h1>

        {/* Todo Input Form */}
        <form onSubmit={handleSubmit(onSubmit)} className='space-y-4 mb-6'>
          <div>
            <input
              {...register('title', { required: true })}
              className='w-full p-2 border rounded'
              placeholder='Enter a new task title'
            />
          </div>
          <div>
            <textarea
              {...register('content')}
              className='w-full p-2 border rounded'
              placeholder='Optional content'
            />
          </div>
          <button
            type='submit'
            className='bg-blue-500 text-white px-4 py-2 rounded'
          >
            {isEditing ? 'Update' : 'Add'}
          </button>
        </form>

        {/* Todo List */}
        <ul className='space-y-3'>
          {tasks.map((task, index) => (
            <li
              key={task.id}
              className={`flex justify-between items-center p-2 bg-gray-100 rounded border ${
                task.done ? 'bg-green-200' : ''
              }`}
            >
              <div>
                <span
                  className={`block font-bold ${
                    task.done ? 'line-through' : ''
                  }`}
                >
                  {task.title}
                </span>
                {task.content && <pre className='text-sm'>{task.content}</pre>}
              </div>
              <div className='space-x-2'>
                <button
                  className='bg-yellow-400 px-2 py-1 rounded text-white'
                  onClick={() => handleEditTask(index)}
                >
                  Edit
                </button>
                <button
                  className='bg-green-500 px-2 py-1 rounded text-white'
                  onClick={() => handleMoveToDone(index)}
                >
                  Done
                </button>
                <button
                  className='bg-red-500 px-2 py-1 rounded text-white'
                  onClick={() => handleMoveToArchive(index)}
                >
                  Archive
                </button>
              </div>
            </li>
          ))}
        </ul>
      </div>
    </div>
  );
};

export default TodoApp;
