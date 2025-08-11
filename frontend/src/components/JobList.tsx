import {useEffect, useState} from 'react';
import { AllCommunityModule, ModuleRegistry } from 'ag-grid-community'; 
import { AgGridReact } from 'ag-grid-react';
import type { ColDef, CellValueChangedEvent}  from 'ag-grid-community';
import 'ag-grid-community/styles/ag-grid.css';
import 'ag-grid-community/styles/ag-theme-alpine.css';

// Register all Community features TODO: Select specific modules used to reduce bundle size
ModuleRegistry.registerModules([AllCommunityModule]);

interface Job {
    id: number; 
    company: string; 
    title: string; 
    status: string; 
    dateApplied: string;
}

export default function JobList() {

    const baseUrl = import.meta.env.VITE_API_BASE_URL
    const [jobs, setJobs] = useState<Job[]>([]);
    const [loading, setLoading] = useState(true); 

    const colDefs: ColDef<Job>[] = [
        { field: "company", headerName: "Company", editable: true },
        { field: "title", headerName: "Title" , editable: true},
        { field: "status", headerName: "Status" , editable: true},
        { field: "dateApplied", headerName: "Date Applied", editable: true}
    ];

    useEffect(() => {
        fetch(`${baseUrl}/jobs`)
            .then((res) => res.json())
            .then((data) => {
                const camelCaseJobs = data.map((job: any) => ({
                    id: job.id,
                    company: job.company,
                    title: job.title,
                    status: job.status,
                    dateApplied: job.date_applied, // map here
                }));
                setJobs(camelCaseJobs); 
                setLoading(false); 
            })
            .catch((err) => {
                console.error("Error fetching jobs:", err);
                setLoading(false); 
            });
    }, []);

    const onCellValueChanged = (event: CellValueChangedEvent) => {
        const updatedJob = event.data as Job;

        fetch(`${baseUrl}/jobs/${updatedJob.id}`, {
            method: "PUT", 
            headers: {
                "Content-Type": "application/json"
            }, 
            body: JSON.stringify({
                company: updatedJob.company, 
                title: updatedJob.title, 
                status: updatedJob.status, 
                date_applied: updatedJob.dateApplied,

            }),
        })
        .then(res => {
            if (!res.ok) {
                throw new Error("Failed to update job");
            }
            return res.json()
        })
        .then(data => {
            console.log("Job updated:", data);
        })
        .catch(err => {
            console.error("Error updating job:", err)
        });
    };

    if (loading) return <p>Loading jobs...</p>; // TODO: Skeleton loading


    return ( 
        <div 
            style={{ 
                maxWidth: 1000, 
                margin: "0 auto",
            }}
        >
            <h2>Job Tracker</h2>
            <div 
                className="ag-theme-alpine"
            >
                <AgGridReact
                    rowData = {jobs}
                    columnDefs={colDefs}
                    defaultColDef={{ flex: 1 }}
                    onCellValueChanged={onCellValueChanged}
                    domLayout="autoHeight" // TODO: Adjust with pagination later
                    theme="legacy"
                />
            </div>

        </div>

     );
}
 
