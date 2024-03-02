export interface Month {
  month: string;
  days: Day[];
}

export interface Day {
  day: number;
  gastos: number;
  ingresos: number;
}
