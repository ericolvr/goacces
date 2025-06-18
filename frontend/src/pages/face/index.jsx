import { useLocation } from 'react-router-dom';


export function Face() {
    const location = useLocation();
    const document = location.state?.document;
  
    return (
      <div>
        <h1>Face Page</h1>
        <p>Document recebido: {document}</p>
      </div>
    );
}