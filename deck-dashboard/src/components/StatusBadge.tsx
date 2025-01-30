// import React from 'react';

interface StatusBadgeProps {
  status: string;
  type?: 'pod' | 'deployment';
}

export default function StatusBadge({ status, type = 'pod' }: StatusBadgeProps) {
  const getStatusColor = () => {
    const colors = {
      Running: 'bg-green-100 text-green-800',
      Available: 'bg-green-100 text-green-800',
      Pending: 'bg-yellow-100 text-yellow-800',
      Progressing: 'bg-blue-100 text-blue-800',
      Failed: 'bg-red-100 text-red-800',
      Unknown: 'bg-gray-100 text-gray-800',
    };
    return colors[status as keyof typeof colors] || colors.Unknown;
  };

  return (
    <span className={`inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium ${getStatusColor()}`}>
      {status}
    </span>
  );
}