import { createGlobalStyle } from "styled-components";

const px2vh = (size, height = 1440) => `${(size / height) * 100}vh`;

const GlobalStyle = createGlobalStyle`
  html {
    scroll-behavior: smooth;
  }
  body {
    margin: 0;
    padding: 0;
    width: 100%;
    overflow-x:hidden;
    font-family: sans-serif;
    -webkit-box-sizing: border-box; /* Safari/Chrome, other WebKit */
    -moz-box-sizing: border-box;    /* Firefox, other Gecko */
    box-sizing: border-box;
  }
  :root {
      font-size: ${px2vh(30)};
    }
    h1{
      font-family: 'Poppins';
      font-weight: bold;
      color: white;
    }
    h2{
      font-family: 'Poppins';
      font-weight: bold;
      color: white;
    }
    h3{
      margin: 0;
      font-family: 'Poppins';
      font-weight: bold;
      color: white;
    }

    a{
      color: inherit; 
    }
    li{
      list-style-type: none;
      margin:0;
    }
    ul{
      list-style-type: none;
      margin:0;
      padding:0;
    }
`;

export default GlobalStyle;
