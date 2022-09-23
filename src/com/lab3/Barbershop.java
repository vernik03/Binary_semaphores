package com.lab3;

import java.util.ArrayList;
import java.util.LinkedList;
import java.util.Queue;
import java.util.Random;
import java.util.concurrent.Semaphore;

import static java.lang.Thread.sleep;

class  i{
    public static int i=0;
}

public class Barbershop {

    public static Integer queue = 0;

    public Queue<Thread> clients;
    public Barber barber;
    private Semaphore barber_sem;
    private Semaphore client_sem;
    private Semaphore client_queue_sem;

    Barbershop() {
        barber_sem = new Semaphore(0);
        client_sem = new Semaphore(0);
        client_queue_sem = new Semaphore(1);
        barber = new Barber(barber_sem, client_sem, client_queue_sem);
        new Thread(barber).start();

        for (int i = 0; i < 10; i++) {
            Client client = new Client(barber_sem, client_sem, client_queue_sem);
            new Thread(client).start();
            queue++;
            System.out.println(queue + " клієнтів в черзі 🟦");
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
    private Semaphore client_queue_sem;
    Barber(Semaphore barber_sem,Semaphore client_sem, Semaphore barber_works) {
        this.barber_sem=barber_sem;
        this.client_sem=client_sem;
        this.client_queue_sem=barber_works;
    }
    public void run(){
        while (true) {
            try {
                barber_sem.acquire();
                hairCutting();
                Barbershop.queue--;
                System.out.println(Barbershop.queue + " клієнтів в черзі 🟥");
                client_sem.release();
                client_queue_sem.release();
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }
    }

    private void hairCutting() throws InterruptedException {

        System.out.println ("Б: Починає стригти");
        try {
            sleep(1000);
        } catch (InterruptedException e) {
            e.printStackTrace();
        }
        System.out.println ("Б: Постриг");
    }

}


class Client implements Runnable{

    public Semaphore barber_sem;
    public Semaphore client_sem;
    public Semaphore client_queue_sem;

    Client(Semaphore barber_sem,Semaphore client_sem, Semaphore barber_works) {
        this.barber_sem=barber_sem;
        this.client_sem=client_sem;
        this.client_queue_sem=barber_works;
    }

    public void run(){
        try {
            client_queue_sem.acquire();
            barber_sem.release();
            System.out.println ("К: Сів у крісло");
            client_sem.acquire();
            System.out.println ("К: Постригся і йде");
        } catch (InterruptedException e) {
            e.printStackTrace();
        }
    }
}


