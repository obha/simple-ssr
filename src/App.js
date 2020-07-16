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
          <Route path="/" render={Home} exact />
          <Route path="/users" component={Users} />
          <Route path="/about" component={About} />
        </Switch>
      </div>
    );
  };

const Home = () =>{
  return <h1>Home</h1>
}

class Users extends React.Component{
  state = {
    
  }
  render(){
    return (
      <div>
        <h1>User</h1>
        <button>login</button>
      </div>
    )
  } 
}

const About = () =>{
    return <h1>About</h1>
}