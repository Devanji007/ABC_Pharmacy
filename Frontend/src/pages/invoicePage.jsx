import { DataGrid } from '@mui/x-data-grid';
import { useEffect, useState } from 'react';
import { toast } from 'react-toastify';
import api from "../api";

const InvoicePage = () => {

    const [rows, setRows] = useState([]);

   

    const getInvoices = async() => {
        try {
            let {data} =  await api.get('/invoices');
           console.log(data)
            setRows(data.data)
           
        } catch (error) {
            toast.error(error);
           
        }
      
    }

    useEffect(() => {
       // getItems();
        getInvoices();
    }, []);


    const columns = [
        {
            field:'invid',
            headerName:'Invoice Id',
            width:90,
            editable: true,
        },

        {
            field:'cusname',
            headerName:'Customer Name',
            width:150,
            editable: true,
        },
        {
            field:'mobno',
            headerName:'Mobile No',
            width:150,
            editable: true,
        },
        {
            field:'email',
            headerName:'Email',
            width:150,
            editable: true,
        },

        {
            field:'address',
            headerName:'Address',
            width:150,
            editable: true,
        },

        {
            field:'billingtype',
            headerName:'Billing Type',
            width:150,
            editable: true,
        },

        {
            field:'totalamount',
            headerName:'Total Amount',
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

    const getRowId = (row) => row.invid

    return(
        <>
            <DataGrid columns={columns} rows={rows} getRowId={getRowId} />
        </>
    );
}

export default InvoicePage;