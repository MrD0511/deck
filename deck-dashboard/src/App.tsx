import { useState } from 'react';
import { 
  Server, 
  Box, 
  Cpu, 
  Database,
} from 'lucide-react';
import Sidebar from './components/Sidebar';
import MetricsCard from './components/MatricsCard';
import SearchBar from './components/SearchBar';
import StatusBadge from './components/StatusBadge';
import { 
  mockMetrics, 
  mockPods, 
  mockServices, 
  mockDeployments 
} from './data/mockData';

function App() {
  const [sidebarOpen, setSidebarOpen] = useState(false);
  const [searchTerm, setSearchTerm] = useState('');

  const filterData = (data: any[]) => {
    return data.filter(item => 
      item.name.toLowerCase().includes(searchTerm.toLowerCase()) ||
      item.namespace.toLowerCase().includes(searchTerm.toLowerCase())
    );
  };

  return (
    <div className="flex min-h-screen bg-gray-50">
      <Sidebar isOpen={sidebarOpen} onToggle={() => setSidebarOpen(!sidebarOpen)} />
      
      <main className={`flex-1 transition-all duration-300 ${sidebarOpen ? 'lg:ml-64' : 'lg:ml-20'}`}>
        <div className="p-6">
          <div className="max-w-7xl mx-auto space-y-8">
            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
              <MetricsCard 
                title="Total Nodes" 
                value={mockMetrics.nodes} 
                icon={<Server className="text-blue-600" size={24} />} 
              />
              <MetricsCard 
                title="Active Pods" 
                value={mockMetrics.pods} 
                icon={<Box className="text-green-600" size={24} />} 
              />
              <MetricsCard 
                title="CPU Usage" 
                value={mockMetrics.cpuUsage} 
                unit="%" 
                icon={<Cpu className="text-purple-600" size={24} />} 
              />
              <MetricsCard 
                title="Memory Usage" 
                value={mockMetrics.memoryUsage} 
                unit="%" 
                icon={<Database className="text-orange-600" size={24} />} 
              />
            </div>

            <SearchBar 
              value={searchTerm} 
              onChange={setSearchTerm} 
              placeholder="Search resources..." 
            />

            <div className="space-y-6">
              {/* Pods Section */}
              <section className="bg-white rounded-lg shadow-sm overflow-hidden">
                <div className="px-6 py-4 border-b border-gray-200">
                  <h2 className="text-lg font-medium text-gray-900">Running Pods</h2>
                </div>
                <div className="overflow-x-auto">
                  <table className="min-w-full divide-y divide-gray-200">
                    <thead className="bg-gray-50">
                      <tr>
                        <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Name</th>
                        <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Namespace</th>
                        <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Status</th>
                        <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">CPU/Memory</th>
                      </tr>
                    </thead>
                    <tbody className="bg-white divide-y divide-gray-200">
                      {filterData(mockPods).map((pod) => (
                        <tr key={pod.name} className="hover:bg-gray-50">
                          <td className="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">{pod.name}</td>
                          <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{pod.namespace}</td>
                          <td className="px-6 py-4 whitespace-nowrap">
                            <StatusBadge status={pod.status} type="pod" />
                          </td>
                          <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                            {pod.cpu} / {pod.memory}
                          </td>
                        </tr>
                      ))}
                    </tbody>
                  </table>
                </div>
              </section>

              {/* Services Section */}
              <section className="bg-white rounded-lg shadow-sm overflow-hidden">
                <div className="px-6 py-4 border-b border-gray-200">
                  <h2 className="text-lg font-medium text-gray-900">Active Services</h2>
                </div>
                <div className="overflow-x-auto">
                  <table className="min-w-full divide-y divide-gray-200">
                    <thead className="bg-gray-50">
                      <tr>
                        <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Name</th>
                        <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Type</th>
                        <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Cluster IP</th>
                        <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Ports</th>
                      </tr>
                    </thead>
                    <tbody className="bg-white divide-y divide-gray-200">
                      {filterData(mockServices).map((service) => (
                        <tr key={service.name} className="hover:bg-gray-50">
                          <td className="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">{service.name}</td>
                          <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{service.type}</td>
                          <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{service.clusterIP}</td>
                          <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                            {service.ports.map((port: any, i: any) => (
                              <div key={i}>
                                {port.port}:{port.targetPort}/{port.protocol}
                              </div>
                            ))}
                          </td>
                        </tr>
                      ))}
                    </tbody>
                  </table>
                </div>
              </section>

              {/* Deployments Section */}
              <section className="bg-white rounded-lg shadow-sm overflow-hidden">
                <div className="px-6 py-4 border-b border-gray-200">
                  <h2 className="text-lg font-medium text-gray-900">Current Deployments</h2>
                </div>
                <div className="overflow-x-auto">
                  <table className="min-w-full divide-y divide-gray-200">
                    <thead className="bg-gray-50">
                      <tr>
                        <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Name</th>
                        <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Replicas</th>
                        <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Status</th>
                        <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Image</th>
                      </tr>
                    </thead>
                    <tbody className="bg-white divide-y divide-gray-200">
                      {filterData(mockDeployments).map((deployment) => (
                        <tr key={deployment.name} className="hover:bg-gray-50">
                          <td className="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">{deployment.name}</td>
                          <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                            {deployment.replicas.current}/{deployment.replicas.desired}
                          </td>
                          <td className="px-6 py-4 whitespace-nowrap">
                            <StatusBadge status={deployment.status} type="deployment" />
                          </td>
                          <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{deployment.image}</td>
                        </tr>
                      ))}
                    </tbody>
                  </table>
                </div>
              </section>
            </div>
          </div>
        </div>
      </main>
    </div>
  );
}

export default App;