import { Link } from "@remix-run/react";
import { Heart, LogOut } from "lucide-react";

const PublicHeader = () => {

    return (
        <header className="bg-white shadow-md">
            <nav className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
                <div className="flex justify-between h-16">
                    {/* Logo and brand */}
                    <div className="flex items-center">
                        <Link to="/" className="flex items-center space-x-2">
                            <Heart className="h-8 w-8 text-rose-500" />
                            <span className="text-2xl font-bold text-gray-900">Matcha</span>
                        </Link>
                    </div>


                    {/* Navigation Links */}
                    <div className="hidden sm:flex sm:items-center sm:space-x-8">
                        <Link
                            to="/login"
                            className="text-gray-600 hover:text-gray-900 px-3 py-2 rounded-md text-sm font-medium"
                        >
                            Login
                        </Link>
                        <Link
                            to="/signup"
                            className="bg-rose-500 text-white hover:bg-rose-600 px-4 py-2 rounded-md text-sm font-medium"
                        >
                            Sign Up
                        </Link>
                    </div>

                    {/* Mobile menu button */}
                    <div className="flex items-center sm:hidden">
                        <button
                            type="button"
                            className="inline-flex items-center justify-center p-2 rounded-md text-gray-400 hover:text-gray-500 hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-inset focus:ring-rose-500"
                            aria-controls="mobile-menu"
                            aria-expanded="false"
                        >
                            <span className="sr-only">Open main menu</span>
                            <svg
                                className="block h-6 w-6"
                                xmlns="http://www.w3.org/2000/svg"
                                fill="none"
                                viewBox="0 0 24 24"
                                stroke="currentColor"
                                aria-hidden="true"
                            >
                                <path
                                    strokeLinecap="round"
                                    strokeLinejoin="round"
                                    strokeWidth={2}
                                    d="M4 6h16M4 12h16M4 18h16"
                                />
                            </svg>
                        </button>
                    </div>
                </div>

                {/* Mobile menu */}
                <div className="sm:hidden" id="mobile-menu">
                    <div className="pt-2 pb-3 space-y-1">
                        <Link
                            to="/about"
                            className="text-gray-600 hover:text-gray-900 block px-3 py-2 rounded-md text-base font-medium"
                        >
                            About
                        </Link>
                        <Link
                            to="/login"
                            className="text-gray-600 hover:text-gray-900 block px-3 py-2 rounded-md text-base font-medium"
                        >
                            Login
                        </Link>
                        <Link
                            to="/signup"
                            className="bg-rose-500 text-white block px-3 py-2 rounded-md text-base font-medium"
                        >
                            Sign Up
                        </Link>
                    </div>
                </div>
            </nav>
        </header>
    );
};

export default PublicHeader;