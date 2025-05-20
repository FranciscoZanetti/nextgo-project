import Task from '@/components/task';
import React from 'react';

export const dynamic = 'force-dynamic';

async function loadTasks() {
  try {
    const baseUrl = process.env.INTERNAL_API_URL;
    
    const res = await fetch(`${baseUrl}/tasks`, {
      next: { revalidate: 10 },
    });

    if (!res.ok) {
      throw new Error(`HTTP error! Status: ${res.status}`);
    }

    const data = await res.json();
    return Array.isArray(data) ? data : [];
  } catch (error) {
    console.error("Error fetching tasks:", error);
    return [];
  }
}

export default async function HomePage() {
  const tasks = await loadTasks();
  myUndefinedFunction();

  return (
    <section className='container mx-auto'>
      <div className='grid grid-cols-3 gap-3 mt-10'>
        {tasks.map((task) => (
          <Task key={task.id} task={task} />
        ))}
      </div>
    </section>
  );
}
