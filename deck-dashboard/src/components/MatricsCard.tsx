import React from 'react';

interface MetricsCardProps {
  title: string;
  value: number;
  unit?: string;
  icon: React.ReactNode;
}

export default function MetricsCard({ title, value, unit, icon }: MetricsCardProps) {
  return (
    <div className="bg-white rounded-lg shadow-sm p-6">
      <div className="flex items-center justify-between">
        <div>
          <p className="text-sm font-medium text-gray-600">{title}</p>
          <p className="mt-2 text-3xl font-semibold text-gray-900">
            {value}
            {unit && <span className="text-lg ml-1">{unit}</span>}
          </p>
        </div>
        <div className="text-blue-600">
          {icon}
        </div>
      </div>
    </div>
  );
}