'use client';

import React, { useState } from 'react';
import Box from '@mui/material/Box';
import TextField from '@mui/material/TextField';
import Button from '@mui/material/Button';
import Typography from '@mui/material/Typography';

import { Login } from './login';
import { useRouter } from 'next/navigation';
import Link from 'next/link';

const LoginPage: React.FC = () => {
    console.log("client")

    const formRefEmail = React.useRef<HTMLInputElement>(null);
    const formRefPassword = React.useRef<HTMLInputElement>(null);
    const router = useRouter();

    const handleLogin = async (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        const formData = new FormData();

        if (formRefEmail.current) formData.append("user_id", formRefEmail.current.value);
        if (formRefPassword.current) formData.append("password", formRefPassword.current.value);

        await Login(formData);

        router.push('/users');
    };

    return (
        <Box
            display="flex"
            flexDirection="column"
            alignItems="center"
            justifyContent="center"
            minHeight="100vh"
        >
            <Typography variant="h4" component="h1" gutterBottom>
                ログイン
            </Typography>
            <Box component="form" onSubmit={handleLogin} sx={{ width: '300px' }}>
                <TextField
                    label="メールアドレス"
                    variant="outlined"
                    fullWidth
                    margin="normal"
                    inputRef={formRefEmail}
                />
                <TextField
                    label="パスワード"
                    type="password"
                    variant="outlined"
                    fullWidth
                    margin="normal"
                    inputRef={formRefPassword}
                />
                <Button type="submit" variant="contained" color="primary" fullWidth>
                    ログイン
                </Button>
            </Box>
            <Box mt={2}>
                <Link href="/signup">
                    <Button variant="outlined" color="secondary" fullWidth>
                        サインアップ
                    </Button>
                </Link>
            </Box>
        </Box>
    );
};

export default LoginPage;