import React from 'react';
import { Link } from 'react-router-dom';

function Navbar() {
  return (
    <nav className="bg-gray-800 text-white p-4">
      <div className="container mx-auto flex justify-between">
        <Link to="/" className="font-bold text-xl hover:text-gray-300">WalletApp</Link>
        <div>
          <Link to="/login" className="mr-4 hover:text-gray-300">Login</Link>
          <Link to="/register" className="hover:text-gray-300">Register</Link>
          {/* More links will be added here once auth is in place, e.g., Dashboard, Logout */}
        </div>
      </div>
    </nav>
  );
}

export default Navbar;
