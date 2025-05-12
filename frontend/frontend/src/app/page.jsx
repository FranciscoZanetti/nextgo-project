import Task from '@/components/task';
import React from 'react'

async function loadTasks() {
  try {
    const res = await fetch("https://nextjs-prisma-crud-nine-woad.vercel.app/api/tasks",  {     next: { revalidate: 10 },   });
    // o dynamic: "force-dynamic" o export const dynamic = 'force-dynamic'
    if (!res.ok) {
      throw new Error(`HTTP error! Status: ${res.status}`);
    }
    const data = await res.json();
    console.log("API Response:", data); // Agregamos esta línea
    return Array.isArray(data) ? data : [];
  } catch (error) {
    console.error("Error fetching tasks:", error);
    return [];
  }
}


export default async function HomePage() {
  const tasks = await loadTasks();
  console.log("tasks: ", tasks);

  return (
    <section className='container mx-auto'>
      <div className='grid grid-cols-3 gap-3 mt-10'>
        {tasks && tasks.map((task) => (
          <Task key={task.id} task={task}/>
        ))}
      </div>
    </section>
  )
}
