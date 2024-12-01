// react-dropzoneを使用したCSVファイルの読み込み
// 'use client';

// import axios from "axios";
// import React, { useState, CSSProperties } from "react";
// import { useDropzone } from "react-dropzone";
// import Dialog from '@mui/material/Dialog';
// import DialogActions from '@mui/material/DialogActions';
// import DialogContent from '@mui/material/DialogContent';
// import DialogContentText from '@mui/material/DialogContentText';
// import DialogTitle from '@mui/material/DialogTitle';

// const CSVReader = () => {
//     console.log("client")
//     const [file, setFile] = useState<File | null>(null);
//     const [open, setOpen] = useState(false);

//     const handleClickOpen = () => {
//         setOpen(true);
//     }

//     const handleClose = () => {
//         setFile(null);
//         setOpen(false);
//     }

//     const handleImport = async () => {
//         if (!file) {
//             return;
//         }
//         const formData = new FormData();
//         formData.append("file", file);
//         try {
//             const res = await axios.post("http://localhost:1323/csv", formData, {
//                 headers: {
//                     "Content-Type": "multipart/form-data"
//                 },
//             });
//             console.log(res);
//             handleClose();
//         } catch (error) {
//             console.error(error);
//         }
//     }

//     const onDrop = (acceptedFiles: File[]) => {
//         const file = acceptedFiles[0];
//         console.log(file);
//         if (file) {
//             setFile(file);
//         }
//     }

//     const { getRootProps, getInputProps, isDragActive, fileRejections } = useDropzone({ 
//         onDrop, 
//         accept: {
//             'test/csv': ['.csv'],
//         },
//     });

//     const fileRejectionsItems = fileRejections.map(({ file, errors}) => (
//         <>
//             <div key={file.name}>
//                 {errors.map((e) => (
//                     <p key={e.code}>
//                         {file.name}は許可された拡張子ではありません。
//                     </p>
//                 ))}
//             </div>
//         </>
//     ));

//     const dropzoneStyle: CSSProperties = {
//         border: '2px dashed #ccc',
//         padding: '20px',
//         textAlign: 'center',
//         cursor: 'pointer',
//         borderColor: isDragActive ? '#000' : '#ccc' // ドラッグ中は黒色に変更
//     };

//     const placeholderStyle: CSSProperties = {
//         color: '#ccc' // 文字を薄くする
//     };

//     return (
//         <div>
//             <h1>CSV Reader</h1>
//             <div>
//                 <button onClick={handleClickOpen}>CSVインポート</button>
//                 <Dialog
//                     open={open}
//                     onClose={handleClose}
//                 >
//                     <DialogTitle>CSVインポート</DialogTitle>
//                     <DialogContent>
//                         <DialogContentText>
//                             インポートするCSVファイルを選択してください。
//                         </DialogContentText>
//                         <div
//                             {...getRootProps({ style: dropzoneStyle})}
//                         >
//                             <input {...getInputProps()} />
//                             {file ? (
//                                 <div>
//                                     <p>ファイル名: {file.name}</p>
//                                     <p>ファイルサイズ: {file.size} bytes</p>
//                                 </div>
//                             ) : (
//                                 <p style={placeholderStyle}>
//                                     ここにファイルをドラッグ＆ドロップ、<br />
//                                     またはクリックでファイルを選択
//                                 </p>
//                             )}
//                         </div>
//                         <DialogActions>
//                             <button onClick={handleImport}>インポート</button>
//                             <button onClick={handleClose}>キャンセル</button>
//                         </DialogActions>
//                         {fileRejectionsItems}
//                     </DialogContent>
//                 </Dialog>
//             </div>
//         </div>
//     )
// }

// export default CSVReader;

// react-papaparseを使用したCSVファイルの読み込み
'use client';

import axios from "axios";
import React, { useState, CSSProperties } from "react";
import Dialog from '@mui/material/Dialog';
import DialogActions from '@mui/material/DialogActions';
import DialogContent from '@mui/material/DialogContent';
import DialogTitle from '@mui/material/DialogTitle';
import {
    useCSVReader,
    lightenDarkenColor,
    formatFileSize,
  } from 'react-papaparse';

const GREY = '#CCC';
const GREY_LIGHT = 'rgba(255, 255, 255, 0.4)';
const DEFAULT_REMOVE_HOVER_COLOR = '#A01919';
const REMOVE_HOVER_COLOR_LIGHT = lightenDarkenColor(
  DEFAULT_REMOVE_HOVER_COLOR,
  40
);
const GREY_DIM = '#686868';

const styles = {
  zone: {
    alignItems: 'center',
    borderWidth: 2,
    borderStyle: 'dashed',
    borderColor: GREY,
    borderRadius: 20,
    display: 'flex',
    flexDirection: 'column',
    height: '100%',
    justifyContent: 'center',
    padding: 20,
  } as CSSProperties,
  file: {
    background: 'linear-gradient(to bottom, #EEE, #DDD)',
    borderRadius: 20,
    display: 'flex',
    height: 120,
    width: 120,
    position: 'relative',
    zIndex: 10,
    flexDirection: 'column',
    justifyContent: 'center',
  } as CSSProperties,
  info: {
    alignItems: 'center',
    display: 'flex',
    flexDirection: 'column',
    paddingLeft: 10,
    paddingRight: 10,
  } as CSSProperties,
  size: {
    backgroundColor: GREY_LIGHT,
    borderRadius: 3,
    marginBottom: '0.5em',
    justifyContent: 'center',
    display: 'flex',
  } as CSSProperties,
  name: {
    backgroundColor: GREY_LIGHT,
    borderRadius: 3,
    fontSize: 12,
    marginBottom: '0.5em',
  } as CSSProperties,
  progressBar: {
    bottom: 14,
    position: 'absolute',
    width: '100%',
    paddingLeft: 10,
    paddingRight: 10,
  } as CSSProperties,
  zoneHover: {
    borderColor: GREY_DIM,
  } as CSSProperties,
  default: {
    borderColor: GREY,
  } as CSSProperties,
  remove: {
    height: 23,
    position: 'absolute',
    right: 6,
    top: 6,
    width: 23,
  } as CSSProperties,
};

const CSVReader = () => {
    console.log("client")
    const [csvFile, setCsvFile] = useState<File | null>(null);
    const [open, setOpen] = useState(false);

    const { CSVReader } = useCSVReader();
    const [zoneHover, setZoneHover] = useState(false);
    const [removeHoverColor, setRemoveHoverColor] = useState(
      DEFAULT_REMOVE_HOVER_COLOR
    );

    const handleClickOpen = () => {
        setOpen(true);
    }

    const handleClose = () => {
        setCsvFile(null);
        setOpen(false);
    }

    const handleImport = async () => {
        if (!csvFile) {
            return;
        }
        const formData = new FormData();
        formData.append("file", csvFile);

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
                        <CSVReader
                            onUploadAccepted={(_result: any, file: any) => {
                                console.log('---------------------------');
                                console.log(file)
                                console.log('---------------------------');
                                setCsvFile(file)
                                setZoneHover(false);
                            }}
                            onDragOver={(event: DragEvent) => {
                                event.preventDefault();
                                setZoneHover(true);
                            }}
                            onDragLeave={(event: DragEvent) => {
                                event.preventDefault();
                                setZoneHover(false);
                            }}
                            >
                            {({
                                getRootProps,
                                acceptedFile,
                                ProgressBar,
                                getRemoveFileProps,
                                Remove,
                            }: any) => (
                                <>
                                <div
                                    {...getRootProps()}
                                    style={Object.assign(
                                    {},
                                    styles.zone,
                                    zoneHover && styles.zoneHover
                                    )}
                                >
                                    {acceptedFile ? (
                                    <>
                                        <div style={styles.file}>
                                        <div style={styles.info}>
                                            <span style={styles.size}>
                                            {formatFileSize(acceptedFile.size)}
                                            </span>
                                            <span style={styles.name}>{acceptedFile.name}</span>
                                        </div>
                                        <div style={styles.progressBar}>
                                            <ProgressBar />
                                        </div>
                                        <div
                                            {...getRemoveFileProps()}
                                            style={styles.remove}
                                            onMouseOver={(event: Event) => {
                                            event.preventDefault();
                                            setRemoveHoverColor(REMOVE_HOVER_COLOR_LIGHT);
                                            }}
                                            onMouseOut={(event: Event) => {
                                            event.preventDefault();
                                            setRemoveHoverColor(DEFAULT_REMOVE_HOVER_COLOR);
                                            }}
                                        >
                                            <Remove color={removeHoverColor} />
                                        </div>
                                        </div>
                                    </>
                                    ) : (
                                    'Drop CSV file here or click to upload'
                                    )}
                                </div>
                                </>
                            )}
                        </CSVReader>
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

export default CSVReader;