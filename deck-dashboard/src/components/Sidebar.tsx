// import React from 'react';
import { 
  LayoutDashboard, 
  Box, 
  Share2, 
  Layers, 
  Settings,
  Menu
} from 'lucide-react';

interface SidebarProps {
  isOpen: boolean;
  onToggle: () => void;
}

export default function Sidebar({ isOpen, onToggle }: SidebarProps) {
  const menuItems = [
    { icon: LayoutDashboard, label: 'Overview', active: true },
    { icon: Box, label: 'Pods' },
    { icon: Share2, label: 'Services' },
    { icon: Layers, label: 'Deployments' },
    { icon: Settings, label: 'Settings' },
  ];

  return (
    <>
      <button
        onClick={onToggle}
        className="lg:hidden fixed top-4 left-4 z-50 p-2 bg-white rounded-md shadow-md"
      >
      <Menu size={24} />
      </button>
      <div className={`
        fixed top-0 left-0 h-full bg-white shadow-lg transition-transform duration-300 z-40
        ${isOpen ? 'translate-x-0 max-h' : '-translate-x-full lg:translate-x-0'}
        w-64 lg:static
      `}>
        <div className="p-4 border-b">
          <h1 className="text-xl font-bold text-gray-800">k3s Dashboard</h1>
        </div>
        <nav className="p-4">
          {menuItems.map((item) => (
            <a
              key={item.label}
              href="#"
              className={`
                flex items-center space-x-3 p-3 rounded-lg mb-2
                ${item.active 
                  ? 'bg-blue-50 text-blue-600' 
                  : 'text-gray-600 hover:bg-gray-50'}
              `}
            >
              <item.icon size={20} />
              <span>{item.label}</span>
            </a>
          ))}
        </nav>
      </div>
    </>
  );
}