'use client';

import React from 'react';
import Box from '@mui/material/Box';
import TextField from '@mui/material/TextField';
import Button from '@mui/material/Button';
import Typography from '@mui/material/Typography';
import { Signup } from './signup';
import { useRouter } from 'next/navigation';

const SignupPage: React.FC = () => {
    const formRefUserID = React.useRef<HTMLInputElement>(null);
    const formRefEmail = React.useRef<HTMLInputElement>(null);
    const formRefPassword = React.useRef<HTMLInputElement>(null);
    const formRefPasswordConfirm = React.useRef<HTMLInputElement>(null);
    const formRefJob = React.useRef<HTMLInputElement>(null);

    const router = useRouter();

    const handleSignup = async (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        const formData = new FormData();
        if (formRefPassword.current && formRefPasswordConfirm.current && formRefPassword.current.value !== formRefPasswordConfirm.current.value) {
            alert('パスワードが一致しません');
            return;
        }

        if (formRefUserID.current) formData.append("user_id", formRefUserID.current.value);
        if (formRefJob.current) formData.append("job", formRefJob.current.value);
        if (formRefEmail.current) formData.append("email", formRefEmail.current.value);
        if (formRefPassword.current) formData.append("password", formRefPassword.current.value);

        await Signup(formData);
        router.push('/login');
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
                サインアップ
            </Typography>
            <Box component="form" onSubmit={handleSignup} sx={{ width: '300px' }}>
                <TextField
                    label="ユーザID"
                    variant="outlined"
                    fullWidth
                    margin="normal"
                    inputRef={formRefUserID}
                />
                <TextField
                    label="職業"
                    variant="outlined"
                    fullWidth
                    margin="normal"
                    inputRef={formRefJob}
                />
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
                <TextField
                    label="パスワード確認"
                    type="password"
                    variant="outlined"
                    fullWidth
                    margin="normal"
                    inputRef={formRefPasswordConfirm}
                />
                <Button type="submit" variant="contained" color="primary" fullWidth>
                    サインアップ
                </Button>
            </Box>
        </Box>
    );
};

export default SignupPage;