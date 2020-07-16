import React from 'react';
import { NavLink, Route, Switch } from 'react-router-dom';

export default props => {
    return (
      <div>
        <ul>
          <li>
            <NavLink to="/">Home</NavLink>
          </li>
          <li>
            <NavLink to="/users">Users</NavLink>
          </li>
          <li>
            <NavLink to="/about">About</NavLink>
          </li>
        </ul>
  
        <Switch>
          <Route
            exact
            path="/"
            render={Home}
          />
          <Route path="/users" component={Users} />
          <Route path="/about" component={About} />
        </Switch>
      </div>
    );
  };

const Home = () =>{
  return <h1>Home</h1>
}

const Users = ()=>{
    return <h1>User</h1>
}

const About = () =>{
    return <h1>About</h1>
}