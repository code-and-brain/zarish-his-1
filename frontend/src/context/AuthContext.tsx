import { createContext, ReactNode, useContext, useState } from 'react';

interface User {
  id: number;
  name: string;
  role: string;
}

interface AuthContextType {
  user: User | null;
  login: (user: User) => void;
  logout: () => void;
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export const AuthProvider = ({ children }: { children: ReactNode }) => {
  // Mock initial user for development
  const [user, setUser] = useState<User | null>({
    id: 1,
    name: 'Dr. Ariful',
    role: 'admin',
  });

  const login = (newUser: User) => setUser(newUser);
  const logout = () => setUser(null);

  return (
    <AuthContext.Provider value={{ user, login, logout }}>
      {children}
    </AuthContext.Provider>
  );
};

export const useAuth = () => {
  const context = useContext(AuthContext);
  if (context === undefined) {
    throw new Error('useAuth must be used within an AuthProvider');
  }
  return context;
};
