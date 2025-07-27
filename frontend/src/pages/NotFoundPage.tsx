
const NotFound = () => {
  return (
    <div className="flex flex-col items-center justify-center text-center px-4 mt-50">
      <h1 className="text-6xl font-extrabold text-stone-100 mb-4">
        404 - Page Not Found
      </h1>
      <p className="text-2xl text-stone-600 dark:text-stone-300 mb-8">
        Sorry, I don&apos;t know how you got here, but the page does not exist.
      </p>
    </div>
  );
};

export default NotFound;
