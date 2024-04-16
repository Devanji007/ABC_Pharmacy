import { useEffect, useState } from "react";
import api from "../api";
import { toast } from "react-toastify";
import { Button } from "@mui/material";
import { useNavigate } from "react-router-dom";

const HomePage = () =>{
    const [count, setCount] = useState(10);

    const navigate = useNavigate()
    
    return(
        <>
           <h1>Home Page</h1>
           <Button onClick={() => navigate('/invoice')}>Go to Invoices</Button>
        </>
    )
}

export default HomePage;