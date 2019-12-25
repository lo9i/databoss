
export class Candidate {
  id: number;
  firstName: string;
  lastName: string;

  userId: string;
  male: boolean;
  phoneNumber: string;

  street: string;
  number: number;
  floor: number;

  department: string;
  city: string;
  zipCode: string;
  state: string;

  birthDate: Date;
}

export class Job {
  id: number;
  name: string;
  from: Date;
  to: Date;
}

