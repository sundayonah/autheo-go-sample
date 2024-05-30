import { useEffect, useState } from "react";
import Modal from 'react-modal';

interface Post {
  createdAt: string;
  title: string;
  body: string;
  id: string;
}

function App() {
  const [posts, setPosts] = useState<Post[]>([]);
  const [isLoading, setIsLoading] = useState(true); // To manage loading state
  const [error, setError] = useState(null);
  const [newPost, setNewPost] = useState({ title: '', body: '' }); // State for form inputs
  const [modalIsOpen, setModalIsOpen] = useState(false);
   const [editPost, setEditPost] = useState<Post | null>(null);

  useEffect(() => {
    const getPosts = async () => {
      try {
        const response = await fetch('http://localhost:8080/api/go-test');
        if (!response.ok) throw new Error('Network response was not ok');
        const data = await response.json();

        // console.log(data);
        setPosts(data);
      } catch (error) {
        console.error("Failed to fetch posts:", error);
        // setError(error?.message);
      } finally {
        setIsLoading(false);
      }
    };
    getPosts();
  }, []);

  // console.log(posts)


 const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    try {
      const response = await fetch('http://localhost:8080/api/create-post', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ title: newPost.title, body: newPost.body }),
      });

      if (!response.ok) throw new Error('Network response was not ok');

      const result = await response.json();
      console.log(result);

      // Optionally, refresh posts or handle success state
      setNewPost({ title: '', body: '' });
      setIsLoading(true);
      const updatedPosts = await fetch('http://localhost:8080/api/go-test').then((res) => res.json());
      console.log(updatedPosts)
      setPosts(updatedPosts);
      setIsLoading(false);
    } catch (error) {
      console.error("Failed to create post:", error);
      setError(error.message);
    }
 };
  
   const handleDelete = async (id: string) => {
    try {
      const response = await fetch(`http://localhost:8080/api/delete-post?id=${id}`, {
        method: 'DELETE',
      });
      console.log(id)

      if (!response.ok) throw new Error('Network response was not ok');

      setPosts(posts.filter(post => post.id !== id));
    } catch (error) {
      console.error("Failed to delete post:", error);
      // setError(error.message);
    }
   };
   const handleEditClick = (post: Post) => {
    setEditPost(post);
    setModalIsOpen(true);
  };

  const handleModalClose = () => {
    setModalIsOpen(false);
  };

  const handleUpdate = async () => {
    try {

    if (!editPost) {
      throw new Error('No post selected for update.');
    }

      const response = await fetch(`http://localhost:8080/api/update-post?id=${editPost.id}`, {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(editPost),
      });

      if (!response.ok) throw new Error('Network response was not ok');

      const updatedPost = await response.json();
      console.log(updatedPost);

      const updatedPosts = posts.map((post) => {
        if (post.id === updatedPost.id) {
          return updatedPost;
        }
        return post;
      });

      setPosts(updatedPosts);
      setModalIsOpen(false);
    } catch (error) {
      console.error("Failed to update post:", error);
      // setError(error.message);
    }
  };

  
  // if posts is null display no post yet
  
  if (posts.length === 0) { 

    return (
      <main className="mt-32">
        <h1 className="text-3xl font-bold text-center">No Post Yet</h1>
      </main>
    )
  }

  return (
    <>
      <div className="">
        <main className="mt-32">
          <h1 className="text-3xl font-bold text-center">Go With REST API</h1>
          <form onSubmit={handleSubmit} className="flex flex-col gap-5 max-w-xl mx-auto p-5">
            <input
              type="text"
              value={newPost.title}
              onChange={(e) => setNewPost({...newPost, title: e.target.value })}
              className="border-2 border-gray-300 p-2 w-full rounded-md"
              placeholder="title"
            />
            <textarea
              value={newPost.body}
              onChange={(e) => setNewPost({...newPost, body: e.target.value })}
              className="border-2 border-gray-300 p-2 w-full rounded-md"
              placeholder="body"
            />
            <input
              type="submit"
              value="Post"
              className="bg-blue-500 text-white p-2 w-full rounded-md cursor-pointer"
            />
          </form>
        </main>
      {/* Displaying posts */}
        {isLoading && <p>Loading...</p>}
        {error && <p>Error: {error}</p>}
    {!isLoading &&!error && ( <div className="max-w-5xl mx-auto grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-4 p-4">
    {posts.map((post) => (
      <article key={post?.id} className="bg-white shadow-md rounded-lg p-4">
        <h2 className="text-xl text-black font-semibold">{post?.title}</h2>
        <p>{post?.body}</p>
           <button
                onClick={() => handleEditClick(post)}
                className="bg-yellow-500 text-white p-2 rounded-md mr-2"
              >
                Edit
              </button>
         <button
                onClick={() => handleDelete(post.id)}
                className="bg-red-500 text-white p-2 rounded-md"
              >
                Delete
              </button>
      </article>
    ))}
  </div>
)}

<Modal
        isOpen={modalIsOpen}
        onRequestClose={handleModalClose}
        contentLabel="Edit Post Modal"
      >
        <h2>Edit Post</h2>
        <form onSubmit={handleUpdate} className="max-w-2xl mx-auto" >
          <input
            type="text"
            value={editPost?.title}
          onChange={(e) => setEditPost({
          ...editPost,
            title: e.target.value,
            createdAt: editPost?.createdAt || "",
            body: editPost?.body || "",
            id: editPost?.id || ""
          })}
              className="border-2 border-gray-300 p-2 w-full rounded-md"
            placeholder="Title"
          />
          <textarea
            value={editPost?.body}
            onChange={(e) => setEditPost({
             ...editPost,
            title: e.target.value,
            createdAt: editPost?.createdAt || "",
            body: editPost?.body || "",
            id: editPost?.id || ""
          })}
            className="border-2 border-gray-300 p-2 w-full rounded-md"
            placeholder="Body"
          />
          <input
            type="submit"
            value="Update"
            className="bg-blue-500 text-white p-2 rounded-md cursor-pointer"
          />
          <button
            onClick={handleModalClose}
            className="bg-gray-500 text-white p-2 rounded-md cursor-pointer ml-2"
          >
            Cancel
          </button>
        </form>
      </Modal>
      </div>
    </>
  );
}

export default App;
