package com.lab3;

import java.util.ArrayList;
import java.util.LinkedList;
import java.util.Queue;
import java.util.Random;
import java.util.concurrent.Semaphore;

import static java.lang.Thread.sleep;

public class Barbershop {

    public Barber barber;
    private Semaphore barber_sem;
    private Semaphore client_sem;
    private Semaphore barber_works;

    Barbershop() {
        barber_sem = new Semaphore(0);
        client_sem = new Semaphore(0);
        barber_works = new Semaphore(0);
        barber = new Barber(barber_sem, client_sem, barber_works);
        new Thread(barber).start();

//        Client client = new Client(barber_sem, client_sem);
//        new Thread(client).start();

        for (int i = 0; i < 100; i++) {
            Client client = new Client(barber_sem, client_sem, barber_works);
            new Thread(client).start();
            Random rand = new Random();
            int int_random = rand.nextInt(1000)+500;
            try {
                sleep(int_random);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }
    }
}

class Barber implements Runnable{

    private Semaphore barber_sem;
    private Semaphore client_sem;
    private Semaphore barber_works;
    Barber(Semaphore barber_sem,Semaphore client_sem, Semaphore barber_works) {
        this.barber_sem=barber_sem;
        this.client_sem=client_sem;
        this.barber_works=barber_works;
    }
    public void run(){
        while (true) {
            try {

                barber_sem.acquire();
                    client_sem.release();
                barber_sem.acquire();
                hairCutting();
                client_sem.release();
//                if (clients.size() > 0) {
//                    clients.element().client_queue_sem.release();
//                }
//                if (clients.size() <= 1) {
//                    clients.notify();
//                }
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }
    }

    private void hairCutting() throws InterruptedException {

        barber_works.release();
        System.out.println ("Б: Начинает стричь");
        try {
            sleep(1000);
        } catch (InterruptedException e) {
            e.printStackTrace();
        }
        System.out.println ("Б: Постриг");
        barber_works.acquire();
    }

}


class Client implements Runnable{

    public Semaphore barber_sem;
    public Semaphore client_sem;
    public Semaphore barber_works;

    Client(Semaphore barber_sem,Semaphore client_sem, Semaphore barber_works) {
        this.barber_sem=barber_sem;
        this.client_sem=client_sem;
        this.barber_works=barber_works;
    }

    public void run(){
        try {
            if (barber_works.availablePermits() == 0) {
                barber_sem.release();
            }
            client_sem.acquire();

//            System.out.println ("К: Пришёл клиент и ждёт");
//            if (client_queue_sem.availablePermits() < 0) {
//                System.out.println (" Пришёл клиент и ждёт");
//                wait();
//                System.out.println (" Дождался");
//            }
//            else {
//                System.out.println (" Пришёл клиент");
//            }
            barber_sem.release();
            System.out.println ("К: Сел в кресло");
            client_sem.acquire();
            System.out.println ("К: Постригся и уходит");
        } catch (InterruptedException e) {
            e.printStackTrace();
        }
        /*if (clients.size() == 1) {
            try {
                sem.acquire();
                System.out.println(" садится в кресло");
                wait();
                System.out.println("постригся");
                sem.release();

            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }*/
    }
}


