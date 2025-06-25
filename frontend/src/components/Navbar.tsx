export default function Navbar() {
  const menuItems = [
    { name: "Rooms", href: "/" },
    { name: "General", href: "/general" },
    { name: "Add New", href: "/add-new" },
  ];

  return (
    <nav className="p-4 fixed top-0 mt-10 ml-5">
      <div className="flex items-center space-x-15">
        {menuItems.map((item) => (
          <a
            key={item.name}
            href={item.href}
            className="text-gray-100 text-xl font-medium transform hover:scale-115"
          >
            {item.name}
          </a>
        ))}
      </div>
    </nav>
  );
}

