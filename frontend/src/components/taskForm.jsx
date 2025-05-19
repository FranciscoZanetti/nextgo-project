"use client";
import { useEffect, useState } from "react";
import { useRouter } from "next/navigation";

export default function taskForm({ taskId }) {
  const router = useRouter();
  const [title, setTitle] = useState("");
  const [description, setDescription] = useState("");
  const API_URL = process.env.NEXT_PUBLIC_API_URL;
  console.log(API_URL);
  

  useEffect(() => {
    if (taskId) {
      fetch(`${API_URL}/tasks/${taskId}`)
        .then(res => res.json())
        .then(data => {
          setTitle(data.title);
          setDescription(data.description);
        })
        .catch(console.error);
    }
  }, [taskId]);

  const onSubmit = async (e) => {
    e.preventDefault();
    const payload = { title, description };

    try {
      const res = await fetch(`${API_URL}/tasks${taskId ? `/${taskId}` : ""}`, {
        method: taskId ? "PUT" : "POST",
        body: JSON.stringify(payload),
        headers: { "Content-Type": "application/json" }
      });

      if (res.ok) {
        router.push("/");
        router.refresh();
      } else {
        console.error("Error:", res.statusText);
      }
    } catch (err) {
      console.error("Error submitting form:", err);
    }
  };

  const handleDelete = async () => {
    try {
      const res = await fetch(`${API_URL}/tasks/${taskId}`, {
        method: "DELETE",
      });
      if (res.ok) {
        router.push("/");
        router.refresh();
      } else {
        console.error("Delete failed:", res.statusText);
      }
    } catch (err) {
      console.error("Error deleting task:", err);
    }
  };

  return (
    <div className='h-screen flex justify-center'>
      <form className='bg-slate-800 p-10 w-1/4' onSubmit={onSubmit}>
        <label htmlFor="title" className='font-bold text-sm'>Título</label>
        <input
          id="title"
          type="text"
          className='border border-gray-400 p-2 mb-4 w-full text-black'
          value={title}
          onChange={(e) => setTitle(e.target.value)}
        />

        <label htmlFor="description" className='font-bold text-sm'>Descripción</label>
        <textarea
          id="description"
          rows="3"
          className='text-black border border-gray-400 p-2 mb-4 w-full'
          value={description}
          onChange={(e) => setDescription(e.target.value)}
        />

        <div className='flex justify-between'>
          <button type="submit" className='bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded'>
            {taskId ? "Actualizar" : "Crear"}
          </button>
          {taskId && (
            <button type="button" className='bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-4 rounded'
              onClick={handleDelete}>Eliminar</button>
          )}
        </div>
      </form>
    </div>
  );
}
