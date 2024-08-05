const express = require('express');
const bodyParser = require('body-parser');
const processController = require('./controllers/processController'); // Ruta correcta

const app = express();
app.use(bodyParser.json());

// Definir la ruta POST para /process
app.post('/process', processController);

const PORT = process.env.PORT || 3000;
app.listen(PORT, () => {
  console.log(`Server running on port ${PORT}`);
});
