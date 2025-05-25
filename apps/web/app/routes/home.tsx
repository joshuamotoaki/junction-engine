import type { Route } from "./+types/home";

// eslint-disable-next-line no-empty-pattern
export function meta({}: Route.MetaArgs) {
    return [
        { title: "New React Router App" },
        { name: "description", content: "Welcome to React Router!" },
    ];
}

export default function Home() {
    return (
        <div className="flex flex-col items-center justify-center h-screen">
            <p className="text-gray-500 mt-4">
                This is a simple React Router application.
            </p>
        </div>
    );
}
