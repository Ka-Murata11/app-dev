'use client';

import axios from "axios";
import React, { useRef, useState } from "react";
import Dialog from '@mui/material/Dialog';
import DialogActions from '@mui/material/DialogActions';
import DialogContent from '@mui/material/DialogContent';
import DialogContentText from '@mui/material/DialogContentText';
import DialogTitle from '@mui/material/DialogTitle';

const CSVReader = () => {
    console.log("client")
    const [file, setFile] = useState<File | null>(null);
    const [open, setOpen] = useState(false);
    const fileInputRef = useRef<HTMLInputElement>(null);

    const handleClickOpen = () => {
        setOpen(true);
    }

    const handleClose = () => {
        setFile(null);
        setOpen(false);
    }

    const handleFileImport = (event: React.ChangeEvent<HTMLInputElement>) => {
        const file = event.target.files?.item(0);
        console.log(file);
        if (file) {
            setFile(file);
        }
    }

    const handleImport = async () => {
        if (!file) {
            return;
        }
        const formData = new FormData();
        formData.append("file", file);
        try {
            const res = await axios.post("http://localhost:1323/csv", formData, {
                headers: {
                    "Content-Type": "multipart/form-data"
                },
            });
            console.log(res);
            handleClose();
        } catch (error) {
            console.error(error);
        }
    }

    const handleDrop = (event: React.DragEvent<HTMLDivElement>) => {
        event.preventDefault();
        event.stopPropagation();
        const file = event.dataTransfer.files?.item(0);
        console.log(file);
        if (file) {
            setFile(file);
        }
    }

    const handleDragOver = (event: React.DragEvent<HTMLDivElement>) => {
        event.preventDefault();
        event.stopPropagation();
    }

    const handleFileSelectClick = () => {
        fileInputRef.current?.click();
    }

    return (
        <div>
            <h1>CSV Reader</h1>
            <div>
                <button onClick={handleClickOpen}>CSVインポート</button>
                <Dialog
                    open={open}
                    onClose={handleClose}
                >
                    <DialogTitle>CSVインポート</DialogTitle>
                    <DialogContent>
                        <DialogContentText>
                            インポートするCSVファイルを選択してください。
                        </DialogContentText>
                        <div
                            onDrop={handleDrop}
                            onDragOver={handleDragOver}
                            onClick={handleFileSelectClick}
                            style={{ border: '2px dashed #ccc', padding: '20px', textAlign: 'center' }}
                        >
                            {file ? (
                                <div>
                                    <p>ファイル名: {file.name}</p>
                                    <p>ファイルサイズ: {file.size} bytes</p>
                                </div>
                            ) : (
                                "ここにファイルをドラッグ＆ドロップ"
                            )}
                        </div>
                        <input
                            type="file"
                            accept=".csv"
                            onChange={handleFileImport}
                            ref={fileInputRef}
                            style={{ display: 'none' }}
                        />
                        <DialogActions>
                            <button onClick={handleImport}>インポート</button>
                            <button onClick={handleClose}>キャンセル</button>
                        </DialogActions>
                    </DialogContent>
                </Dialog>
            </div>
        </div>
    )
}
