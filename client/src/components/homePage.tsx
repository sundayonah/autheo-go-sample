export default function Home() {
    return (
        <main className="mt-32">
            <h1 className="text-3xl font-bold text-center">Go With REST API</h1>
            <form action="" className="flex flex-col gap-5 max-w-xl mx-auto p-5">
                <input type="text" className="border-2 border-gray-300 p-2 w-full rounded-md" placeholder="Title" />
                <input type="text" className="border-2 border-gray-300 p-2 w-full rounded-md" placeholder="Body" />
                <input type="submit" className="bg-blue-500 text-white p-2 w-full rounded-md" value="Post" />
            </form>
        </main>
    );
}