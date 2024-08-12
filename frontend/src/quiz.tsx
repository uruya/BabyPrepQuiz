import React, { useEffect, useState} from "react";

type Question = {
  id:number;
  question:string;
  options:string[];
  answer:number;
}

const Quiz: React.FC = () => {
  const [questions, setQuestions] = useState<Question[]>([]);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    fetch("http://localhost:8080/api/questions")
      .then((response) => {
        if (!response.ok) {
          throw new Error("Network response was not ok");
        }
        return response.json();
      })
      .then((data) => {
        console.log("Fetched data: ", data);

        if (!Array.isArray(data)) {
          throw new Error("Expected data to be an array");
        }

        const formattedData = data.map((item: any) => ({
          id: item.ID,
          question: item.Question,
          options: item.Options,
          answer: item.Answer,
        }));
        setQuestions(formattedData);
      })
      .catch((error) => {
        console.error("There was a problem with the fetch operation:", error);
        setError(error.message);
      });
  }, []);

  if (error) {
    return <div>Error: {error}</div>;
  }

  if(!questions) {
    return <div>Loading...</div>;
  }

  return (
    <div>
      {questions.length === 0 ? (
        <div>Loading...</div>
      ) : (
        questions.map((q, index) => (
          <div key={index}>
            <h3>{q.question}</h3>
            {q.options.map((option, i) => (
              <button key={i}>{option}</button>
            ))}
          </div>
        ))
      )}
    </div>
  );
};

export default Quiz;