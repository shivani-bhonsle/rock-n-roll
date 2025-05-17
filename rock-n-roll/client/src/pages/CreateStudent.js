import * as React from "react";

const CreateStudent = () => {
  const [input, setInput] = React.useState({});
  const [users, setUsers] = React.useState([]);

  const eventHandler = (e) => {
    let name = e.target.name;
    let value = e.target.value;
    setInput((prev) => {
      return { ...prev, [name]: value };
    });
  };

  const clickHandler = (e) => {
    e.preventDefault();
    console.log(input);
    const student = {
      Name: input.name,
      Email: input.email,
    };
    console.log(student);
    const url = `https://jsonplaceholder.typicode.com/users`;
    fetch(url)
      .then((res) => res.json())
      .then((data) => {
        console.log(data);
        const names = data.map((d) => d.name);
        console.log(names);
        setUsers(names);
      });
  };

  return (
    <>
      <div>
        <form>
          <label for="name">Student Name</label>
          <input
            id="name"
            type="text"
            name="name"
            value={input.name || ""}
            onChange={eventHandler}
          />
          <br />
          <label for="email">Email</label>
          <input
            id="email"
            type="email"
            name="email"
            value={input.email || ""}
            onChange={eventHandler}
          />
          <br />
          <button onClick={clickHandler}>Submit</button>
        </form>
        {users.length > 0 && (
          <select name="trick">
            Select a random name
            {users.map((user) => (
              <option value={user}>{user}</option>
            ))}
          </select>
        )}
      </div>
      <div></div>
    </>
  );
};

export default CreateStudent;
