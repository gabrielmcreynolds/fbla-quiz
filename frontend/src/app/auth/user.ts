import * as moment from 'moment/moment';

export class User {
  id: string;
  email: string;
  name: string;
  testsTaken: number;
  totalScores: number;
  totalTime: moment.Duration;
}
