'use client';

import React from "react";
import Box from '@mui/material/Box';
import TextField from '@mui/material/TextField';
import Button from '@mui/material/Button';
import MenuItem from '@mui/material/MenuItem';
import { CreateUser } from "./fetcher";

const CreateUserForm = () => {
    const formRefName = React.useRef<HTMLInputElement>(null);
    const formRefEmail = React.useRef<HTMLInputElement>(null);
    const formRefRole = React.useRef<HTMLSelectElement>(null); 
    const formRefJob = React.useRef<HTMLInputElement>(null);

    const AddUser = async (event: React.FormEvent<HTMLFormElement>) => {

        event.preventDefault();
        const formData = new FormData();

        if (formRefName.current) formData.append("name", formRefName.current.value);
        if (formRefEmail.current) formData.append("email", formRefEmail.current.value);
        if (formRefRole.current) formData.append("role", formRefRole.current.value);
        if (formRefJob.current) formData.append("job", formRefJob.current.value);

        await CreateUser(formData);
    }

    return (
        <>
        <Box component="form" onSubmit={AddUser} >
            <TextField
                label="名前"
                inputRef={formRefName}
                variant="outlined"
                required
            />
            <TextField
                label="メールアドレス"
                type="email"
                inputRef={formRefEmail}
                variant="outlined"
                required
            />
            <TextField
                label="権限"
                select
                inputRef={formRefRole}
                variant="outlined"
                required
                defaultValue="01"
            >
                <MenuItem value="01">Admin</MenuItem>
                <MenuItem value="02">User</MenuItem>
            </TextField>
            <TextField
                label="職業"
                inputRef={formRefJob}
                variant="outlined"
                required
            />
            <Button type="submit" variant="contained" color="primary">
                ユーザーを作成
            </Button>
        </Box>
        </>
    )
}

export default CreateUserForm;