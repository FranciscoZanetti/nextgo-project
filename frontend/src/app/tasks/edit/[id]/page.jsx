import TaskForm from "@/components/taskForm";

export default function EditTaskPage({ params }) {
  return <TaskForm taskId={params.id} />;
}
