// app/components/AppHeader.tsx
import { Link } from "@remix-run/react";
import { Heart, User, Users, MessageCircle, LogOut } from "lucide-react";
import { useFetcher } from "@remix-run/react";


export const AppHeader = () => {
    const fetcher = useFetcher();

    const handleLogout = () => {
        fetcher.submit(null, {
            method: "post",
            action: "/logout"  // 別ルートのactionを指定
        });
    }


    return (
        <header className="bg-white shadow-md">
            <nav className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
                <div className="flex justify-between h-16">
                    <div className="flex items-center space-x-8">
                        <Link to="/app" className="flex items-center space-x-2">
                            <Heart className="h-8 w-8 text-rose-500" />
                            <span className="text-2xl font-bold text-gray-900">Matcha</span>
                        </Link>
                        <div className="hidden sm:flex sm:items-center sm:space-x-4">
                            <Link
                                to="/map"
                                className="flex items-center space-x-1 text-gray-600 hover:text-gray-900 px-3 py-2"
                            >
                                <Users className="w-5 h-5" />
                                <span>Matches</span>
                            </Link>
                            <Link
                                to="/my-profile"
                                className="flex items-center space-x-1 text-gray-600 hover:text-gray-900 px-3 py-2"
                            >
                                <MessageCircle className="w-5 h-5" />
                                <span>Messages</span>
                            </Link>
                            <Link
                                to="/my-profile"
                                className="flex items-center space-x-1 text-gray-600 hover:text-gray-900 px-3 py-2"
                            >
                                <User className="w-5 h-5" />
                                <span>Profile</span>
                            </Link>
                        </div>
                    </div>
                    <div className="flex items-center">
                        <button onClick={handleLogout}
                            type="button"
                            className="flex items-center space-x-1 text-gray-600 hover:text-gray-900 px-3 py-2"
                        >
                            <LogOut className="w-5 h-5" />
                            <span>Logout</span>
                        </button>
                    </div>
                </div>
            </nav>
        </header>
    );
};