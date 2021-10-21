import { createPool, Pool } from 'mysql2/promise';

export async function connect(): Promise<Pool> {
  const connection = await createPool({
    host: 'localhost',
    user: 'root',
    database: 'funsiba',
    connectionLimit: 10,
    password: 'root',
  });
  return connection;
}
