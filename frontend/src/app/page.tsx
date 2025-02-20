'use client';

import { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';

export default function Home() {
  const router = useRouter();
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    try {
      
      if (typeof window !== 'undefined') {
        
      const token = localStorage.getItem('token');
      if (token) {
        router.push('/dashboard');
      } else {
        router.push('/auth/login');
      }
    }
    } catch (error) {
      console.error('Navigation error:', error);
      router.push('/auth/login');
    } finally {
      setIsLoading(false);
    }
  }, [router]);

  if (isLoading) {
    return (
      <div className="flex items-center justify-center min-h-screen">
        <div>Loading...</div>
      </div>
    );
  }

  return null;
}
