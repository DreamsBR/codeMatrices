module.exports = (req, res) => {
    const matrices = req.body;
  
    if (!matrices || !matrices.Q || !matrices.R) {
      return res.status(400).json({ error: 'Matrices invalidas' });
    }
    console.log("matrices : ", matrices);
    
  
    // Funciones para calcular estadísticas
    const calculateStats = (matrix) => {
      let max = -Infinity;
      let min = Infinity;
      let sum = 0;
      let count = 0;
  
      matrix.forEach(row => {
        row.forEach(value => {
          if (typeof value !== 'number') {
            throw new Error('Matrix no contiene numeros');
          }
          max = Math.max(max, value);
          min = Math.min(min, value);
          sum += value;
          count += 1;
        });
      });
  
      const average = sum / count;
      return { max, min, average, total: sum };
    };
  
    // Verificar si la matriz es diagonal
    const isDiagonal = (matrix) => {
      for (let i = 0; i < matrix.length; i++) {
        for (let j = 0; j < matrix[i].length; j++) {
          if (i !== j && matrix[i][j] !== 0) {
            return false;
          }
        }
      }
      return true;
    };
  
    try {
      const qStats = calculateStats(matrices.Q);
      const rStats = calculateStats(matrices.R);
  
      res.json({
        q: {
          ...qStats,
          isDiagonal: isDiagonal(matrices.Q),
        },
        r: {
          ...rStats,
          isDiagonal: isDiagonal(matrices.R),
        }
      });
    } catch (error) {
      res.status(500).json({ error: error.message });
    }
  };
  