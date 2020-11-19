export class Question {
  question: string;
  choices: Array<string> | null;
  correctChoice: (string | boolean);
}
