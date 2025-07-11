import { Link } from 'react-router-dom'

export default function Navbar({ workspaceID }: { workspaceID: string }) {

  const menuItems = [
    { name: "Rooms", href: `/${workspaceID}/rooms` },
    { name: "General", href: `/${workspaceID}/generals/` },
    { name: "Add New", href: `/${workspaceID}/add-new/` },
  ];

  return (
    <nav className="p-4 top-0 mt-1 ml-5">
      <div className="flex items-center space-x-15">
        {menuItems.map((item) => (
          <Link
            key={item.name}
            to={item.href}
            className="text-stone-900 dark:text-stone-100 text-xl font-medium transform hover:scale-115"
          >
            {item.name}
          </Link>
        ))}
      </div>
    </nav>
  );
}

