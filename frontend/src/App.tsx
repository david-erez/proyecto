import { useEffect, useState } from "react";

function App() {
  const [message, setMessage] = useState("Conectando con el backend...");

  useEffect(() => {
    fetch("/api/hello")
      .then((response) => response.json())
      .then((data) => {
        setMessage(data.message);
      })
      .catch((error) => {
        console.error(error);
        setMessage("Error conectando con el backend");
      });
  }, []);

  return (
    <div>
      <h1>Frontend React</h1>
      <p>{message}</p>
    </div>
  );
}

export default App;
