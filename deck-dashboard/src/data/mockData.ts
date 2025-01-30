import { Pod, Service, Deployment, ClusterMetrics } from '../types/kubernates';

export const mockMetrics: ClusterMetrics = {
  nodes: 3,
  pods: 24,
  deployments: 8,
  services: 12,
  cpuUsage: 67,
  memoryUsage: 45,
  storageUsage: 32,
};

export const mockPods: Pod[] = [
  {
    name: 'nginx-deployment-6b474476c4-2j9xk',
    namespace: 'default',
    status: 'Running',
    cpu: '125m',
    memory: '256Mi',
    createdAt: '2024-03-10T15:30:00Z',
  },
  {
    name: 'redis-master-5d9844bf75-q8z7x',
    namespace: 'data',
    status: 'Running',
    cpu: '250m',
    memory: '512Mi',
    createdAt: '2024-03-10T14:20:00Z',
  },
  {
    name: 'prometheus-7bc4b4d9d9-jf2k7',
    namespace: 'monitoring',
    status: 'Running',
    cpu: '500m',
    memory: '1Gi',
    createdAt: '2024-03-10T12:15:00Z',
  },
];

export const mockServices: Service[] = [
  {
    name: 'nginx-service',
    namespace: 'default',
    type: 'LoadBalancer',
    clusterIP: '10.96.0.1',
    ports: [
      { port: 80, targetPort: 8080, protocol: 'TCP' },
    ],
  },
  {
    name: 'redis-master',
    namespace: 'data',
    type: 'ClusterIP',
    clusterIP: '10.96.0.2',
    ports: [
      { port: 6379, targetPort: 6379, protocol: 'TCP' },
    ],
  },
];

export const mockDeployments: Deployment[] = [
  {
    name: 'nginx-deployment',
    namespace: 'default',
    replicas: { desired: 3, current: 3, ready: 3 },
    status: 'Available',
    image: 'nginx:1.21',
    createdAt: '2024-03-10T10:00:00Z',
  },
  {
    name: 'redis-master',
    namespace: 'data',
    replicas: { desired: 1, current: 1, ready: 1 },
    status: 'Available',
    image: 'redis:7.0',
    createdAt: '2024-03-10T09:45:00Z',
  },
];