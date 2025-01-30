// Type definitions for Kubernetes resources
export interface Pod {
    name: string;
    namespace: string;
    status: 'Running' | 'Pending' | 'Failed' | 'Succeeded' | 'Unknown';
    cpu: string;
    memory: string;
    createdAt: string;
  }
  
  export interface Service {
    name: string;
    namespace: string;
    type: 'ClusterIP' | 'NodePort' | 'LoadBalancer';
    clusterIP: string;
    ports: Array<{
      port: number;
      targetPort: number;
      protocol: 'TCP' | 'UDP';
    }>;
  }
  
  export interface Deployment {
    name: string;
    namespace: string;
    replicas: {
      desired: number;
      current: number;
      ready: number;
    };
    status: 'Available' | 'Progressing' | 'Failed';
    image: string;
    createdAt: string;
  }
  
  export interface ClusterMetrics {
    nodes: number;
    pods: number;
    deployments: number;
    services: number;
    cpuUsage: number;
    memoryUsage: number;
    storageUsage: number;
  }