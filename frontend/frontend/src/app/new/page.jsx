"use client"
import React, { useEffect, useState } from 'react'
import { useRouter } from 'next/navigation';

export default function NewPage({params}) {

    const router = useRouter();

    const [title, setTitle] = useState("");
    const [description, setDescription] = useState("");

    useEffect(() => {
        if (params.id){
            fetch(`/api/tasks/${params.id}`)
            .then(res => res.json())
            .then(data => {
                setTitle(data.title);
                setDescription(data.description);
            });
        }
    }, [params.id]);

    const onSubmit = async (e) => {
        e.preventDefault();
        // con el useEffect que actualiza title y description, ya no hacen falta
        //const title = e.target.title.value;
        //const description = e.target.description.value;
        
        if (params.id){
            const res = await fetch(`"/api/tasks/${params.id}`, {
                method: "PUT",
                body: JSON.stringify({title, description}),
                headers:{
                    "Content-Type": "application/json",

                }
            });
            const data = await res.json();
        }
        else{
            const res = await fetch("/api/tasks", {
                method: "POST",
                body: JSON.stringify({title, description}),
                headers: {
                    "Content-Type": "application/json"
                }
            });
            const data = await res.json();
        }

        router.refresh();
        router.push("/");
    }

  return (
    <div className='h-screen flex justify-center'>
        <form action="" className='bg-slate-800 p-10 w-1/4' onSubmit={onSubmit}>
            <label htmlFor="title" className='font-bold text-sm'>Titulo de la tarea</label>
            <input id='title' type="text" className='border border-gray-400 p-2 mb-4 w-full text-black' placeholder='Titulo' value={title} onChange={(e) => setTitle(e.target.value)}/>
            <label htmlFor="description" className='font-bold text-sm'>Descripcion de la tarea</label>
            <textarea name="" id="description" cols="30" rows="3" className='text-black border border-gray-400 p-2 mb-4 w-full' placeholder='descripcion' value={description} onChange={(e) => setDescription(e.target.value)}></textarea>
            <div className='flex justify-between'>
                <button type='submit' className='bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded'>Crear</button>
                {
                    params.id && (
                        <button type='button' className='bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-4 rounded'
                            onClick={async (e) => {
                                const res = await fetch(`api/tasks/${params.id}`, {
                                    method: "DELETE",
                                });
                                const data = await res.json();
                                router.refresh();
                                router.push("/");
                            }}
                        >Eliminar</button>
                    )
                }
            </div>
        </form>
    </div>
  )
}
