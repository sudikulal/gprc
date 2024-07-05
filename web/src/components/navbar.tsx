// components/Navbar.tsx
import Link from 'next/link';
import { useState, ChangeEvent, FormEvent } from 'react';

const Navbar = () => {
  const [searchQuery, setSearchQuery] = useState('');

  const handleSearchChange = (e: ChangeEvent<HTMLInputElement>) => {
    setSearchQuery(e.target.value);
  };

  const handleSearchSubmit = (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    // Implement search logic here
    console.log('Searching for:', searchQuery);
  };

  return (
    <div className="navbar bg-base-100 shadow-lg">
      <div className="flex-1">
        <Link href="/">
          <a className="btn btn-ghost normal-case text-xl">Home</a>
        </Link>
      </div>
      <div className="flex-none gap-2">
        <form onSubmit={handleSearchSubmit} className="form-control">
          <input
            type="text"
            placeholder="Search"
            className="input input-bordered"
            value={searchQuery}
            onChange={handleSearchChange}
          />
        </form>
        <div className="dropdown dropdown-end">
          <label tabIndex={0} className="btn btn-ghost normal-case">Sort</label>
          <ul tabIndex={0} className="dropdown-content menu p-2 shadow bg-base-100 rounded-box w-52">
            <li><a>Newest</a></li>
            <li><a>Oldest</a></li>
            <li><a>Alphabetical</a></li>
          </ul>
        </div>
        <Link href="/contact">
          <a className="btn btn-ghost normal-case">Contact</a>
        </Link>
        <Link href="/logout">
          <a className="btn btn-ghost normal-case">Logout</a>
        </Link>
      </div>
    </div>
  );
};

export default Navbar;
