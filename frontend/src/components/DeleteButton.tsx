export type DeleteButtonProps = {
  handleSubmit: () => void;
}

const DeleteButton = ({ handleSubmit }: DeleteButtonProps) => {
  const onClick = () => {
    if (confirm("Are you sure you want to delete this?")) {
      handleSubmit();
    }
  };

  return (
    <button
      onClick={onClick}
      className="px-4 py-2 bg-red-600 text-white rounded hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-red-500"
    >
      Delete
    </button>
  );
};

export default DeleteButton
