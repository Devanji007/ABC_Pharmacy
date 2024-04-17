import { useEffect, useState } from 'react';
import { toast } from 'react-toastify';
import api from "../api";
import { DataGrid } from '@mui/x-data-grid';



const ItemPage = () =>{

    const [rows, setRows] = useState([]);

     const getItems = async() => {
        try {
            let {data} =  await api.get('/items');
            setRows(data.data)
           
        } catch (error) {
           toast.error(error);
        }
      
    }


    useEffect(() => {
         getItems();
         
     }, []);

     const columns = [
        {
            field:'id',
            headerName:'Item Id',
            width:90,
            editable: true,
        },

        {
            field:'itemname',
            headerName:'Item Name',
            width:150,
            editable: true,
        },
        {
            field:'unitprice',
            headerName:'Unit Price',
            width:150,
            editable: true,
        },
        {
            field:'category',
            headerName:'Category',
            width:150,
            editable: true,
        },

        {
            field:'createdat',
            headerName:'Created At',
            width:150,
            editable: true,
        },

        {
            field:'updatedat',
            headerName:'Updated At',
            width:150,
            editable: true,
        },

        
    ]

    const getRowId = (row) => row.id

    return(
        <>
            <DataGrid columns={columns} rows={rows} getRowId={getRowId} />
        </>
    );
}

export default ItemPage;
